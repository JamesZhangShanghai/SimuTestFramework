package simuctrl

import (
	"container/list"
	"fmt"
	"gofer/lib/timer"
	"time"

	"mavenir.com/nrup/up_sp/impl/services/simutestsvc/simutestserver/simudb"
)

func waitAllFinish(sc *Scenario) {
	ch := make(chan string)
	//start timer
	timeOutHandler := func(seq int32) {
		sc.state = "Error"
		ch <- "testTerm"
	}
	duration := sc.timeOut
	timer := timer.NewTimer()
	timer.Start(duration*1000, timeOutHandler, true)
	//trigger some message

	//wait the result
	//may watch the DB for terminate the test
	//time.Sleep(time.Second * 10)
	tickTimer := time.NewTicker(2 * time.Second)
	fmt.Println("The te*Scenario, st case started.")
	//judge the test result, and save it to DB

WAIT:
	for {
		select {
		case <-ch:
			fmt.Println("TestCase ", sc.name, " wait Timeout")
			break WAIT
		case <-tickTimer.C:
			{
				if simudb.IsAllProcFinished() {
					break
				}
				fmt.Println("Waiting all proc done for scenario ", sc.name)
			}
		}
	}
}

func testE1SetupReq(sc *Scenario) {
	fmt.Println("testE1SetupReq")
	var procList *list.List
	procList = list.New()
	AddProc(procList, "UP-E1-SETUP-REQ", "default", "RESPONSE")
	//save to DB

	waitAllFinish(sc)

	sc.state = "Finished"
}

func testE1Release(sc *Scenario) {
	fmt.Println("testE1Release")
	var procList *list.List
	procList = list.New()
	AddProc(procList, "UP-E1-RELEASE-REQ", "default", "RESPONSE")
	//save to DB

	waitAllFinish(sc)

	sc.state = "Finished"
}

func testE1Reset(sc *Scenario) {
	fmt.Println("testE1ResetReq")
	var procList *list.List
	procList = list.New()
	AddProc(procList, "UP-E1-RESET-REQ", "default", "RESPONSE")
	//save to DB

	waitAllFinish(sc)

	sc.state = "Finished"
}
