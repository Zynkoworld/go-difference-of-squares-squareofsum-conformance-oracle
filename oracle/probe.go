package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



func SquareOfSum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s * s
}

func SumOfSquares(n int) (s int) {
	for i := 1; i <= n; i++ {
		s += i * i
	}
	return s
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []int{1, 5, 100}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, SquareOfSum(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}
