package main

const (
	idCheckTmCost   = 60
	bodyCheckTmCost = 120
	xRayCHeckTmCost = 180
)

func main() {
	// println("--- 运行素数筛示例 ---")
	// RunPrimeSieve()
	// println()

	// println("--- 方案1：顺序设计 ---")
	// runCheckExample()
	// println()

	// println("--- 方案2：并行方案 ---")
	// runCheckExample2()
	// println()

	println("--- 方案3：并发方案 ---")
	runCheckExample3()
	println()

}
