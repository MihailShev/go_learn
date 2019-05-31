package main

import (
	"fmt"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, _ := ntp.Time("ntp5.stratum2.ru")
	fmt.Println(ntpTime.Format("15:04:05"))
}
