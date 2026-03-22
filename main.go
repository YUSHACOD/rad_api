package main

import (
	"fmt"
	"log"
	// "os"

	"rad_api/rad_api"
)

func main() {
	log.Println("testing rad_api")

	var r_api rad_api.RadIpcState
	err := r_api.Init()
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer r_api.Release()

	// r_api.SendCommand(rad_api.RadCmd(os.Args[1]), "")
	// r_api.SendCommand(rad_api.CMD_STEP_INTO, "")
	err = r_api.SendCommand(rad_api.CMD_CLEAR_BREAKPOINTS, "")
	if err != nil {
		fmt.Printf("Send error => %v\n", err)
	}
}
