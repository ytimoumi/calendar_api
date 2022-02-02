package logger

import (
	"fmt"
	"log"
	"runtime/debug"
)

func Log(clientName, clientIp string, err error, withDebug bool) {

	log.Println(err)
	var s = ""
	if withDebug {
		s = string(debug.Stack())
	}
	mail := New(clientName, fmt.Sprintf("=> %s \n=> Client Name : %s \n=> Client IP : %s \n=> %s", err, clientName, clientIp, s))
	if err := mail.Send(); err != nil {
		log.Println(err)
	}

}
