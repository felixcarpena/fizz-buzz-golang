package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Number int
	Value  string
}

func fizzbuzz(rules []Rule, value int) string {
	var result string
	for _, rule := range rules {
		if value%rule.Number == 0 {
			result += rule.Value + " "
		}
	}

	if result != "" {
		return strings.TrimSpace(result)
	} else {
		return fmt.Sprint(value)
	}
}

func configurableFizzbuzz(value int) string {
	rules := []Rule{Rule{3, "fizz"}, Rule{5, "buzz"}, Rule{7, "pop"}}
	return fizzbuzz(rules, value)
}

func main() {
	i, _ := strconv.Atoi(os.Args[1:][0])
	fmt.Printf(configurableFizzbuzz(i))
}
