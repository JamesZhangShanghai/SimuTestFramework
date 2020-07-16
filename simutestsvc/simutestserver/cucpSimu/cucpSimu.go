package cucpSimu

import "fmt"
import "mavenir.com/nrup/up_sp/impl/services/simutestsvc/simutestserver/simuctrl"

type MethodTable struct {
	name     string
	ID       string
	msgType  string
	procFunc func()
}

var methodTab []MethodTable

func init() {
	methodTab = make([]MethodTable, 10)

	tab := MethodTable{
		name:     "test e1 setup",
		ID:       "test1",
		msgType:  "UP-E1-SETUP-REQ",
		procFunc: respToUpE1SetupReq,
	}
	methodTab = append(methodTab, tab)
}

func respToUpE1SetupReq() {
	fmt.Println("simu cp respToUpE1SetupReq")
	//get the expt table
	var behaviour, proc string
	exptProc := GetExptProc("UP-E1-SETUP-REQ")
	if exptProc == nil {
		behaviour = "default"
		proc = "default"
	} else {
		procedure = exptProc.procedure
		behaviour = exptProc.behaviour
	}

	switch behaviour {

	case "RESPONSE":
	{
		fmt.Println("RESPONSE behaviour")
		//send the resp
		break;
	}
    case "FAILURE":
    {
		fmt.Println("FAILURE behaviour")
		//send the failure resp
		break;
	}
	case "MUTE":
	{
		fmt.Println("MUTE behaviour")
		//don't resp
		break;
	}
	default:
	{
		fmt.Println("Unknown behaviour")
	}

	exptProc.state = "finished"
}
