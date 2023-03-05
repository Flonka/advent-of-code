package day11

type Monkey struct {
	items        []int
	operation    func(int) int
	inspectCount int
}

type test struct {
	parameter     int
	successTarget int
	failureTarget int
}
