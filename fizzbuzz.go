package main

import (
	"fmt"
	"os"
	"strconv"
)

type Rule struct {
	Numbers []int
	Value   string
}

func fizzbuzz(rules []Rule, value int) string {
	for _, rule := range rules {
		for _, number := range rule.Numbers {
			if value%number == 0 {
				return rule.Value
			}
		}
	}
	return fmt.Sprint(value)
}

func configurableFizzbuzz(value int) string {
	rules := []Rule{Rule{[]int{15}, "fizz-buzz"}, Rule{[]int{3}, "fizz"}, Rule{[]int{5}, "buzz"}}
	return fizzbuzz(rules, value)
}

func main() {
	i, _ := strconv.Atoi(os.Args[1:][0])
	fmt.Printf(configurableFizzbuzz(i))
}
