package sctpserver



//#cgo LDFLAGS: ${SRCDIR}/sctpserver.a -lsctp
//#include "cwrap.h"
import "C"

import "fmt"

func StartSctpServer() {
	fmt.Println("Start Sctp server")
	C.startSctp()
}