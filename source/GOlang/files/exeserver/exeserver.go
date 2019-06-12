package main

import (
    "context"
    "io"
    "log"
    "net/http"
	"time"
	"sync"
	"os/exec"
	"fmt"
)
var wg sync.WaitGroup

func main() {
    log.Printf("main: starting HTTP server")

	wg.Add(1)
	srv1 := &http.Server{Addr: ":8081"}
	//srv2 := &http.Server{Addr: ":8082"}

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		out, err := exec.Command("./datatypes.exe").Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("COMMAND EXECUTED SUCCESSFULLY")
		//fmt.Println(string(out))
        io.WriteString(w, string(out)) // "hello world\n")
	})

	
    go func() {
        if err := srv1.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatalf("ListenAndServe(): %s", err)
		}
		
		wg.Done()
    }()

    
	
    // go func() {
    //     if err := srv2.ListenAndServe(); err != http.ErrServerClosed {
    //         log.Fatalf("ListenAndServe(): %s", err)
	// 	}
	// 	wg.Done()
    // }()


	log.Printf("main: serving for 10 Seconds")

    time.Sleep(10 * time.Second)

    log.Printf("main: stopping HTTP server")

    // now close the server gracefully ("shutdown")
    // timeout could be given with a proper context
    // (in real world you shouldn't use TODO()).
    if err := srv1.Shutdown(context.TODO()); err != nil {
        panic(err) // failure/timeout shutting down the server gracefully
    }
	// if err := srv2.Shutdown(context.TODO()); err != nil {
    //     panic(err) // failure/timeout shutting down the server gracefully
	// }
	wg.Wait()

    log.Printf("main: done. exiting")
}