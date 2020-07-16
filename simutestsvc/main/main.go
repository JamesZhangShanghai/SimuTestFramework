package main

import (
	"fmt"
	"runtime/debug"
	"time"

	"gcommon/impl/services/common"
	"gcommon/impl/services/common/cimadapter"
	"gcommon/impl/services/common/httpserver"
	"gcommon/impl/services/common/lmaas"

	"up_sp/impl/services/iwfsvc/iwfclient"
	"up_sp/impl/services/simutestsvc/simutestserver"
	"up_sp/impl/services/simutestsvc/simutestserver/sctpserver"
	"up_sp/impl/services/simutestsvc/simutestserver/simuctrl"

	"gofer/lib/err"
	"gofer/lib/log"
	"gofer/lib/timer"
	"gofer/registrar"
)

var (
	logModule int
	counter   int
)

func init() {
	logModule = log.Register("MAIN", 2)
	log.SetLogLevel(log.Inf)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			print(logModule, "goroutine exited because of ", r, string(debug.Stack()))
		}
		cimadapter.UnloadCIMAdapter()
		lmaas.Unload()
		time.Sleep(2 * time.Second)
	}()

	fmt.Println("Test1")
	fmt.Println("Test2")
	fmt.Println("Test3")

	startHTTP()
	simutestserver.Load()
	iwfclient.Load()

	exit := make(chan struct{})
	registrar.Start(exit)

	//do the main func
	var simu_start_wait uint32 = 5
	counter = 0
	timeOutHandler := func(int32) {
		go sctpserver.StartSctpServer()
		go simutestserver.PeriodSendReq()
		simutestserver.InitSimuTestObj()
		go simuctrl.SimuTest(exit)
		for {
			if print := log.Error(logModule); print != nil {
				print(logModule, "started TimeoutHandler")
			}
			print(logModule, "Do something int the simu main routine counter: .... ", counter, "\n")
			time.Sleep(20 * time.Second)
			counter++
		}
	}

	timer := timer.NewTimer()
	if timer == nil {
		if print := log.Error(logModule); print != nil {
			print(logModule, "New timer failed. err = ", err.Errno(err.EOutofMem))
		}
	}
	timer.Start(simu_start_wait*1000, timeOutHandler, true)

	<-exit
	registrar.Stop()
	time.Sleep(time.Second * 3)
}

func startHTTP() {
	if e := httpserver.InitHTTPServer(); e != nil {
		if print := log.Error(logModule); print != nil {
			print(logModule, e)
		}
	}

	if e := cimadapter.LoadCimAdapter(); e != nil {
		if print := log.Error(logModule); print != nil {
			print(logModule, e)
		}
	}

	//init log2server
	if e := lmaas.Load(); e != nil {
		if print := log.Error(logModule); print != nil {
			print(logModule, e)
		}
	}

	//register http route
	if e := common.RegisterHTTPVersion(); e != nil {
		if print := log.Error(logModule); print != nil {
			print(logModule, e)
		}
	}

	//start the server
	httpserver.StartHTTPServer()
}
