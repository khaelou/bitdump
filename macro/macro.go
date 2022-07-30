package macro

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var ProductChannel = make(chan ProductSignal, math.MaxInt8)

type ExecMacro func() map[int]string

type ProductSignal struct {
	Product       interface{}
	WorkerID      int
	WorkerFactory string
	WorkerRole    string
}

func ExecuteMacro(id int, factory, role string, task string, execFunc ExecMacro) {
	execFunc()

	product := execFunc()
	productSignal := ProductSignal{Product: product, WorkerID: id, WorkerFactory: factory, WorkerRole: role}
	ProductChannel <- productSignal
}

func TicketPool() map[int]string {
	var genAmount = 6
	numberMap := make(map[int]string)

	for i := 1; i <= genAmount; i++ {
		if i != genAmount {
			normalNumber := fmt.Sprintf("#%d", genNumber1To70())
			fmt.Println("\t", normalNumber)

			for _, v := range numberMap {
				if normalNumber == v {
					fmt.Println("DUP**", normalNumber)
				}
			}

			numberMap[i] = normalNumber
		} else {
			goldNumber := fmt.Sprintf("#%d", genNumber1To25())
			fmt.Println("\t", goldNumber, "GOLD")

			for _, v := range numberMap {
				if goldNumber == v {
					fmt.Println("DUP**", goldNumber)
				}
			}

			numberMap[i] = goldNumber
		}
	}

	if hasDupes(numberMap) {
		fmt.Println("^ DUPLICATES FOUND.")
	}
	fmt.Println()

	return numberMap
}

func hasDupes(m map[int]string) bool {
	x := make(map[string]struct{})

	for _, v := range m {
		if _, has := x[v]; has {
			return true
		}
		x[v] = struct{}{}
	}

	return false
}

func genNumber1To70() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 70 // Any # between 1-70
	number := rand.Intn(max-min+1) + min

	return number
}

func genNumber1To25() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 25 // Any # between 1-25
	number := rand.Intn(max-min+1) + min

	return number
}
