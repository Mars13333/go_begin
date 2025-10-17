package main

import "time"

func idCheck() int {
	time.Sleep(time.Millisecond * time.Duration(idCheckTmCost))
	println("\tidCheck ok")
	return idCheckTmCost
}

func bodyCheck() int {
	time.Sleep(time.Millisecond * time.Duration(bodyCheckTmCost))
	println("\tbodyCheck ok")
	return bodyCheckTmCost
}

func xRayCheck() int {
	time.Sleep(time.Millisecond * time.Duration(xRayCHeckTmCost))
	println("\txRayCheck ok")
	return xRayCHeckTmCost
}

func airportSecurityCheck() int {
	println("airportSecurityCheck ...")
	total := 0
	total += idCheck()
	total += bodyCheck()
	total += xRayCheck()
	println("airportSecurityCheck ok")
	return total
}

// 方案1：顺序设计
func runCheckExample() {
	total := 0
	passengers := 30
	for i := 0; i < passengers; i++ {
		total += airportSecurityCheck()
	}
	println("total time cost:", total)
	// total time cost: 10800
}
