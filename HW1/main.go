package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("ntp5.stratum2.ru")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ntpTime.Format("15:04:05"))
	}
}