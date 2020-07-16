package simudb

import (
	"container/list"
	"fmt"
)

//ExptProc : the expect msg behaviour
type ExptProc struct {
	msgType   string
	procedure string
	behaviour string
	state     string
}

type TestCase struct {
	name     string
	exptProc *list.List
	state    string
}

var GTestCase *TestCase

func init() {
	GTestCase = &TestCase{
		name:     "",
		exptProc: nil,
		state:    "init",
	}
	GTestCase.exptProc = list.New()
}

//AddExptProc : Add a proc to the simuDB :
func AddExptProc(msg, proc, behr string) {
	exptProc := &ExptProc{
		msgType:   msg,
		procedure: proc,
		behaviour: behr,
	}
	GTestCase.exptProc.PushBack(exptProc)
}

func GetExptProc(msgtype string) *ExptProc {
	var proc ExptProc
	if GTestCase.state == "true" {
		for item := GTestCase.exptProc.Front(); nil != item; item = item.Next() {
			fmt.Println(item.Value)
			proc = item.Value.(ExptProc)
			fmt.Println("The expt proc:", proc)
			if proc.msgType == msgtype {
				return &proc
			}
		}
	}
	return nil
}

func GetAllExptProc() *ExptProc {
	var proc *ExptProc
	if GTestCase.state == "true" {
		for item := GTestCase.exptProc.Front(); nil != item; item = item.Next() {
			fmt.Println(item.Value)
			proc = item.Value.(*ExptProc)
			fmt.Println("The GET expt proc:", proc)
			//if proc.msgType == msgtype {
			//return &proc
			//}
		}
	}
	return nil
}

func IsAllProcFinished() bool {
	var proc *ExptProc
	for item := GTestCase.exptProc.Front(); nil != item; item = item.Next() {
		fmt.Println(item.Value)
		proc = item.Value.(*ExptProc)
		fmt.Println("The expt proc:", proc)
		if proc.state != "finished" {
			return false
		}
	}
	return true
}
