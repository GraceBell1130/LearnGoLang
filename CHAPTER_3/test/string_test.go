package test

import (
	"fmt"
	"strconv"
	"strings"
)

func ExampleHasConsonatSuffix() {
	//fmt.Println(hangul.HasConsonatSuffix("Go 언어"))
	//fmt.Println(hangul.HasConsonatSuffix("그럼"))
	//fmt.Println(hangul.HasConsonatSuffix("우리 밥 먹고 합시다."))
	// Output:
	// false
	// true
	// false
}

func Example_printBytes() {
	s := "가나다"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}
	fmt.Println()
	// Output:
	// eab080eb8298eb8ba4
}

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	fmt.Println(fmt.Sprint(s, "zz"))
	fmt.Println(fmt.Sprintf("%szz", s))
	// Output:
	// abcdef
	// abcdef
	// abcdefzz
	// abcdefzz
}

func Example_StringToInteger() {
	var i int
	var k int64
	var f float64
	var s string
	var err error
	i, err = strconv.Atoi("350")
	fmt.Println(i == 350)
	fmt.Println(err)
	k, err = strconv.ParseInt("cc7fdd", 16, 32)
	fmt.Println(k == 13402077)
	k, err = strconv.ParseInt("0xcc7fdd", 0, 32) // 2번째 아규먼트가 0일 경우 strings의 접두어 값을 보고 판단
	fmt.Println(k == 13402077)
	f, err = strconv.ParseFloat("3.14", 64)
	fmt.Println(f == 3.14)
	s = strconv.Itoa(340)
	fmt.Println(s == "340")
	s = strconv.FormatInt(13402077, 16)
	fmt.Println(s == "cc7fdd")
	// Output:
	// true
	// <nil>
	// true
	// true
	// true
	// true
	// true
}

func Eval(expr string) int {
	var ops []string
	var nums []int
	pop := func() int { // 리터럴 함수
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}

	reduce := func(higher string) { // 리터럴 함수
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if strings.Index(higher, op) < 0 {
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				return
			}

			b, a := pop(), pop()
			switch op {
			case "+":
				nums = append(nums, a+b)
			case "-":
				nums = append(nums, a-b)
			case "*":
				nums = append(nums, a*b)
			case "/":
				nums = append(nums, a/b)
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops = append(ops, token)
		case "+", "-":
			reduce("+-*/")
			ops = append(ops, token)
		case "*", "/":
			reduce("*/")
			ops = append(ops, token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}

	reduce("+-*/")
	return nums[0]
}

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	fmt.Println(Eval("1 - 2 + 3"))
	fmt.Println(Eval("3 * ( 3 + 1 * 3 ) / 2"))
	fmt.Println(Eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))

	// Output:
	// 5
	// 3
	// 2
	// 9
	// 18
}
