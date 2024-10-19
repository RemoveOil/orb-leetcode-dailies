package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	q1405 "solutions/1405-longest-happy-string"
	q1545 "solutions/1545-kth-bit"
	q2044 "solutions/2044-count-max-bitwise-subsets"
	q650 "solutions/650-maximum-swap"
	"solutions/qwerr"
	"strconv"
)

// TODO: Flag for reading custom input, or file.
func main() {
	questionNumber, err := parseQuestionNumber()
	qwerr.AbortWhen(err)

	switch questionNumber {
	case q1405.QUESTION_NO:
		fmt.Println(q1405.LongestHappyString(7, 7, 1))
	case q650.QUESTION_NO:
		fmt.Println(q650.MaximumSwap(2736))
		fmt.Println(q650.MaximumSwap(27376))
		fmt.Println(q650.MaximumSwap(0))
		fmt.Println(q650.MaximumSwap(93))
		fmt.Println(q650.MaximumSwap(4))
	case q2044.QUESTION_NO:
		fmt.Println(q2044.CountMaxOrSubsets([]int{3, 1}))
		fmt.Println(q2044.CountMaxOrSubsets([]int{2, 2, 2}))
		fmt.Println(q2044.CountMaxOrSubsets([]int{3, 2, 1, 5}))
		fmt.Println(q2044.CountMaxOrSubsets([]int{10}))
	case q1545.QUESTION_NO:
		fmt.Println(q1545.FindKthBit(4, 7))
		fmt.Println(q1545.FindKthBit(4, 8))
		fmt.Println(q1545.FindKthBit(4, 9))
		fmt.Println(q1545.FindKthBit(4, 10))
		fmt.Println(q1545.FindKthBit(4, 11))
		fmt.Println(q1545.FindKthBit(4, 12))
		fmt.Println(q1545.FindKthBit(4, 13))
		fmt.Println(q1545.FindKthBit(4, 14))
	default:
		log.Fatal("You fucked up chicken.")
	}

}

func parseQuestionNumber() (int, error) {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return 0, fmt.Errorf("expected exactly 1 argument. Got %v", len(args))
	}

	qNumber, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, errors.Join(err, errors.New("could not parse question number"))
	}

	return qNumber, nil

}
