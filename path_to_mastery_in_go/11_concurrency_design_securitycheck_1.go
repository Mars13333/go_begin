package main

import (
	"fmt"
	"time"
)

/*顺序设计*/

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCheckTMCost = 180
)

func idCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	fmt.Println("\tidCheck ok")
	return idCheckTmCost
}

func bodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	fmt.Println("\tbodyCheck ok")
	return bodyCheckTmCost
}

func xRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(xRayCheckTMCost))
	fmt.Println("\txRayCheck ok")
	return xRayCheckTMCost
}

func airportSecurityCheck() int {
	fmt.Println("airportSecurityCheck...")
	total := 0
	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()
	fmt.Println("airportSecurityCheck ok")
	return total
}

func main() {
	total := 0
	passengers := 30
	for i := 0; i < passengers; i++ {
		total += airportSecurityCheck()
	}
	fmt.Println("total time cost:", total)
}

/*
airportSecurityCheck...
        idCheck ok
        bodyCheck ok
        xRayCheck ok
airportSecurityCheck ok
...
airportSecurityCheck...
        idCheck ok
        bodyCheck ok
        xRayCheck ok
airportSecurityCheck ok
total time cost: 10800
*/
