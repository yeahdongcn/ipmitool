package main

import (
	"fmt"
	"log"

	"github.com/yeahdongcn/ipmitool/pkg/client"
)

func main() {
	cl, err := client.NewClient("10.10.132.1", 0, "admin", "Password@_")
	if err != nil {
		log.Fatal(err)
	}

	status, err := cl.Power.Status()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status)

	if status == client.PowerStateOff {
		err := cl.Power.On()
		if err != nil {
			log.Fatal(err)
		}
	}
}
