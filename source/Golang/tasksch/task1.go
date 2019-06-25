package main

import (
	"fmt"
	"time"

	"github.com/capnspacehook/taskmaster"
	"github.com/rickb777/date/period"
)

func main() {
	var err error
	taskService, err := taskmaster.Connect("", "", "", "")
	if err != nil {
		panic(err)
	}
	defer taskService.Cleanup()
	newTaskDef := taskService.NewTaskDefinition()

	var choice int
	fmt.Println("if you want to create a new task enter '1' if you want to get details about existing task press '2':")
	fmt.Scanln(&choice)
	if choice == 1 {

		var str string
		fmt.Println("enter path:")
		fmt.Scanln(&str)
		newTaskDef.AddExecAction(str, "/c $(Arg0)", "", "basic")
		//newTaskDef.AddExecAction("D:\\go-workspace\\src\\mail\\mail.go", " ", "", "basic")
		var choice2 int
		fmt.Println("if you want to create a repeating daily task press 1 else 2:")
		fmt.Scanln(&choice2)
		if choice2 == 1 {
			var hrs, min, sec int
			fmt.Println("hours:")
			fmt.Scanln(&hrs)
			fmt.Println("min:")
			fmt.Scanln(&min)
			fmt.Println("sec:")
			fmt.Scanln(&sec)
			var (
				defaultTime   time.Time
				defaultPeriod period.Period
				repFive       = period.NewHMS(hrs%25, min%61, sec%61)
			)

			//newTaskDef.AddTimeTrigger(period.NewHMS(0, 0, 0), time.Now().Add(5*time.Second))
			newTaskDef.AddTimeTriggerEx(defaultPeriod, "", time.Now().Add(3*time.Second), defaultTime, defaultPeriod, defaultPeriod, repFive, true, true)

		} else {
			var (
				defaultPeriod period.Period
			)
			var date int
			fmt.Println("enter date")
			fmt.Scanln(&date)
			newTaskDef.AddMonthlyTrigger(18, taskmaster.June, defaultPeriod, time.Now().Add(5*time.Second))
		}
		newTaskDef.RegistrationInfo.Author = "saitej"
		newTaskDef.RegistrationInfo.Description = "just a simple task"
		newTaskDef.RegistrationInfo.Date = time.Time{}
		var str1 string
		fmt.Println("enter task name:")
		fmt.Scanln(&str1)
		newTask, _, err := taskService.CreateTask("\\saitej\\"+str1, newTaskDef, true)
		if err != nil {
			panic(err)
		}
		runningTask, err := newTask.Run([]string{"timeout 42"})
		if err != nil {
			panic(err)
		}
		defer runningTask.Release()

		fmt.Printf("Running task PID: %d\n", runningTask.EnginePID)
	} else {
		var t3 string

		fmt.Println("enter name of task ")
		fmt.Scanln(&t3)
		oldTaskDef, err1 := taskService.GetRegisteredTask("\\" + t3)
		if err1 != nil {
			panic(err)
		}
		fmt.Println("LastRunTime:", oldTaskDef.LastRunTime)
		fmt.Println("lastTaskResult:", oldTaskDef.LastTaskResult)
		fmt.Println("MissedRuns:", oldTaskDef.MissedRuns)
		fmt.Println("NextRunTime:", oldTaskDef.NextRunTime)
		fmt.Println("state of task:", oldTaskDef.State)
		var choice1 int
		fmt.Println("if you want to stop task press 1")
		fmt.Scanln(&choice1)
		if choice1 == 1 {
			oldTaskDef.Stop()
		} else {
			fmt.Println("wrongchoice")
		}

	}
}
