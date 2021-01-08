// package main

// import (
// 	"io/ioutil"
// 	"fmt"
// 	"os"
// 	"log"
// 	"strings"
// )

// func main() {
// 	var fname string = "" //path of file
// 	fmt.Println("Enter name of file: ")
// 	fmt.Scanln(&fname)

// 	fileInfo, stat_err := os.Stat(fname)
// 	if nil != stat_err {
// 		if os.IsNotExist(stat_err) {
// 		log.Fatal("File does not exist: ",stat_err)
// 		}
// 	}
// 	log.Println("File Stat is: ",fileInfo)

// 	//opening file in read_only mode
// 	fptr, open_err := os.OpenFile(fname,os.O_RDONLY,0666)
// 	if nil != open_err {
// 		log.Fatal("File Open Error for read mode: ",open_err)
// 	}
// 	log.Println("File Name after renaming is: ",fptr.Name())

// 	//Reading contents of file
// 	my_content, read_err := ioutil.ReadAll(fptr)
// 	if nil != read_err {
// 		log.Fatal("File Read Error: ", read_err)
// 	}
// 	fmt.Println("File content is:")
// 	fmt.Println(string(my_content))

// 	count:= strings.Count(string(my_content), "are")
// 	fmt.Println(count)
// 	fptr.Close()

// 	if(count > 3) {
// 		fmt.Println("Restarting a service")
// 	}

// }

package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"log"
	"os"
	"strings"

	"path/filepath"
	"golang.org/x/sys/windows/svc/mgr"

	"time"

	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"

	"io/ioutil"
	"syscall"
	"os/exec"
)

func exePath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("%s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("%s is directory", p)
		}
	}
	return "", err
}

func installService(name, desc string) error {
	exepath, err := exePath()
	if err != nil {
		return err
	}
	
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", name)
	}
	s, err = m.CreateService(name, exepath, mgr.Config{DisplayName: desc}, "is", "auto-started")
	if err != nil {
		return err
	}
	defer s.Close()
	err = eventlog.InstallAsEventCreate(name, eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		s.Delete()
		return fmt.Errorf("SetupEventLogSource() failed: %s", err)
	}
	return nil
}

func removeService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("service %s is not installed", name)
	}
	defer s.Close()
	err = s.Delete()
	if err != nil {
		return err
	}
	err = eventlog.Remove(name)
	if err != nil {
		return fmt.Errorf("RemoveEventLogSource() failed: %s", err)
	}
	return nil
}

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop, pause or continue.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func startService(name string) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	err = s.Start("is", "manual-started")
	if err != nil {
		return fmt.Errorf("could not start service: %v", err)
	}
	return nil
}

func controlService(name string, c svc.Cmd, to svc.State) error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		return fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	status, err := s.Control(c)
	if err != nil {
		return fmt.Errorf("could not send control=%d: %v", c, err)
	}
	timeout := time.Now().Add(10 * time.Second)
	for status.State != to {
		if timeout.Before(time.Now()) {
			return fmt.Errorf("timeout waiting for service to go to state=%d", to)
		}
		time.Sleep(300 * time.Millisecond)
		status, err = s.Query()
		if err != nil {
			return fmt.Errorf("could not retrieve service status: %v", err)
		}
	}
	return nil
}

var elog debug.Log

type myservice struct{}

func (m *myservice) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown | svc.AcceptPauseAndContinue
	changes <- svc.Status{State: svc.StartPending}
	fasttick := time.Tick(500 * time.Millisecond)
	slowtick := time.Tick(2 * time.Second)
	tick := fasttick
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
loop:
	for {
		select {
		case <-tick:
			beep()
			elog.Info(1, "beep")
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				// Testing deadlock from https://code.google.com/p/winsvc/issues/detail?id=4
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				// golang.org/x/sys/windows/svc.TestExample is verifying this output.
				testOutput := strings.Join(args, "-")
				testOutput += fmt.Sprintf("-%d", c.Context)
				elog.Info(1, testOutput)
				break loop
			case svc.Pause:
				changes <- svc.Status{State: svc.Paused, Accepts: cmdsAccepted}
				tick = slowtick
			case svc.Continue:
				changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
				tick = fasttick
			default:
				elog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func runService(name string, isDebug bool) {
	var err error
	if isDebug {
		elog = debug.New(name)
	} else {
		elog, err = eventlog.Open(name)
		if err != nil {
			return
		}
	}
	defer elog.Close()

	elog.Info(1, fmt.Sprintf("starting %s service", name))
	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(name, &myservice{})
	if err != nil {
		elog.Error(1, fmt.Sprintf("%s service failed: %v", name, err))
		return
	}
	elog.Info(1, fmt.Sprintf("%s service stopped", name))
}

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)

func beep() {
	beepFunc.Call(0xffffffff)
}

func my_service() {
	const svcName = "AppMgmt" //"C:/Users/Navya/go/src/files/countword/a.exe"

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if !isIntSess {
		runService(svcName, false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		runService(svcName, true)
		return
	case "install":
		err = installService(svcName, "navya4")
	case "remove":
		err = removeService(svcName)
	case "start":
		err = startService(svcName)
	case "stop":
		err = controlService(svcName, svc.Stop, svc.Stopped)
	case "pause":
		err = controlService(svcName, svc.Pause, svc.Paused)
	case "continue":
		err = controlService(svcName, svc.Continue, svc.Running)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, svcName, err)
	}
	return
}

//own functions
func file_exist(fname string) (string) {
	fileInfo, stat_err := os.Stat(fname) 
	if nil != stat_err {
		if os.IsNotExist(stat_err) {
		log.Println("File does not exist: ",stat_err)
		}
	}
	log.Println("File Stat is: ",fileInfo)
	return "exist"
}

func readFile(fname string) (int) {
	//opening file in read_only mode
	fptr, open_err := os.OpenFile(fname, os.O_RDONLY, 0666)
	if nil != open_err {
		log.Fatal("File Open Error for read mode: ", open_err)
	}
	log.Println("File Name is: ", fptr.Name())

	//Reading contents of file
	my_content, read_err := ioutil.ReadAll(fptr)
	if nil != read_err {
		log.Fatal("File Read Error: ", read_err)
	}
	fmt.Println("File content is:")
	fmt.Println(string(my_content))

	count := strings.Count(string(my_content), "are")
	fmt.Println(count)
	fptr.Close()

	return count
}

func main() {
	var fname string = "" //path of file
	fmt.Println("Enter name of file: ")
	fmt.Scanln(&fname)

	val := file_exist(fname)
	if "exist" == val {
		fmt.Println("File exist check done")
	} else {
		log.Fatalln("File exist check fail")
	}
	
	count := readFile(fname)
	if count > 3 {
		fmt.Println("Operating service")
		my_service()
		//fmt.Println("SUCCESS")
		out, err := exec.Command("echo.bat").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(out))
	}

}