package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Day1Solver struct{}

func (d Day1Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	start := time.Now()
	defer close(answerChan)
	defer close(errorChan)
	defer close(doneChan)
	file, err := os.Open(filename)

	if err != nil {
		errorChan <- err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	numbers := []int{}

	for scanner.Scan() {
		num := scanner.Text()

		number, err := strconv.ParseInt(num, 10, 64)
		if err != nil {
			errorChan <- err
			return
		}
		numbers = append(numbers, int(number))
	}

	result1 := 0
	result2 := 0

	for _, base := range numbers {
		for _, right := range numbers {
			if base+right == 2020 {
				result1 = base * right
			}
		}
	}

	for index1, base := range numbers {
		for index2, right := range numbers {
			for index3, third := range numbers {
				if index1 == index2 || index2 == index3 || index3 == index1 {
					break
				}
				if base+right+third == 2020 {
					result2 = base * right * third
				}
			}
		}
	}

	answerChan <- result1
	answerChan <- result2

	time.Sleep(time.Second)

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("\n\n⏱️ Execution took %v time! ⏱️\n\n", elapsed)
	doneChan <- true
}
