package test

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type ExerciseStrSet map[string]struct{}
type ExercisePrecMap map[string]ExerciseStrSet

func ExerciseNewStrSet(strs ...string) ExerciseStrSet {
	m := ExerciseStrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

func ExerciseEval(opMap map[string]BinOp, prec ExercisePrecMap, expr string) (int, error) {
	ops := []string{"("} // 초기 여는 괄호
	var nums []int
	pop := func() int {
		last := nums[len(nums)-1]
		nums = nums[:len(nums)-1]
		return last
	}
	reduce := func(nextOp string) {
		for len(ops) > 0 {
			op := ops[len(ops)-1]
			if _, higher := prec[nextOp][op]; nextOp != ")" && !higher {
				// 더 낮은 순위 연산자이므로 여기서 계산 종료
				return
			}
			ops = ops[:len(ops)-1]
			if op == "(" {
				// 괄호를 제거하였으므로 종료
				return
			}
			b, a := pop(), pop()
			if f := opMap[op]; f != nil {
				nums = append(nums, f(a, b))
			}
		}
	}
	for _, token := range strings.Split(expr, " ") {
		if token == "" {
			continue
		} else if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산을 하고 제거
			reduce(token)
		} else {
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
	}
	reduce(")") // 초기의 여는 괄호까지 모두 계산
	return nums[0], nil
}

func ExerciseNewEvaluator(ooMap map[string]BinOp, prec ExercisePrecMap) func(expr string) int {
	return func(expr string) int {
		result, err := ExerciseEval(ooMap, prec, expr)
		if err == nil {
			return result
		} else {
			fmt.Println("error :", err)
		}
		return 0
	}
}

func ExampleExercise() {
	eval := ExerciseNewEvaluator(map[string]BinOp{
		"**": func(a, b int) int {
			if a == 1 {
				return 1
			}
			if b < 0 {
				return 0
			}
			r := 1
			for i := 0; i < b; i++ {
				r *= a
			}
			return r
		},
		"*":   func(a, b int) int { return a * b },
		"/":   func(a, b int) int { return a / b },
		"mod": func(a, b int) int { return a % b },
		"+":   func(a, b int) int { return a + b },
		"-":   func(a, b int) int { return a - b },
	}, ExercisePrecMap{
		"**":  ExerciseNewStrSet(),
		"*":   ExerciseNewStrSet("**", "*", "/", "mod"),
		"/":   ExerciseNewStrSet("**", "*", "/", "mod"),
		"mod": ExerciseNewStrSet("**", "*", "/", "mod"),
		"+":   ExerciseNewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   ExerciseNewStrSet("**", "*", "/", "mod", "+", "-"),
	})

	fmt.Println(eval("5"))
	fmt.Println(eval("1  +  2"))
	fmt.Println(eval("1  -  2  -  4"))
	fmt.Println(eval("(      3 - 2      ** 3 ) * ( -2 )"))
	fmt.Println(eval("3 * ( 3 + 1 * 3     ) / ( -2 )"))
	fmt.Println(eval("3 * ( ( 3 +    1  ) * 3 ) / 2"))
	fmt.Println(eval("1 + 2 ** 10 *     2"))
	fmt.Println(eval("2 ** 3 mod    3"))
	fmt.Println(eval("2 **      2 ** 3"))
	fmt.Println(eval("2 ** a + 2 ** 3"))

	//Output:
	// 5
	// 3
	// -5
	// 10
	// -9
	// 18
	// 2049
	// 2
	// 256
	// error : strconv.Atoi: parsing "a": invalid syntax
	// 0
}

func EvalReplaceAll(in string) {
	eval := func(expr string) string {
		expr = strings.Trim(expr, "{ }")
		result, err := ExerciseEval(map[string]BinOp{
			"**": func(a, b int) int {
				if a == 1 {
					return 1
				}
				if b < 0 {
					return 0
				}
				r := 1
				for i := 0; i < b; i++ {
					r *= a
				}
				return r
			},
			"*":   func(a, b int) int { return a * b },
			"/":   func(a, b int) int { return a / b },
			"mod": func(a, b int) int { return a % b },
			"+":   func(a, b int) int { return a + b },
			"-":   func(a, b int) int { return a - b },
		}, ExercisePrecMap{
			"**":  ExerciseNewStrSet(),
			"*":   ExerciseNewStrSet("**", "*", "/", "mod"),
			"/":   ExerciseNewStrSet("**", "*", "/", "mod"),
			"mod": ExerciseNewStrSet("**", "*", "/", "mod"),
			"+":   ExerciseNewStrSet("**", "*", "/", "mod", "+", "-"),
			"-":   ExerciseNewStrSet("**", "*", "/", "mod", "+", "-"),
		}, expr)

		if err == nil {
			return strconv.Itoa(result)
		} else {
			fmt.Println("error :", err)
		}
		return ""
	}
	rx := regexp.MustCompile(`{[^}]+}`)
	out := rx.ReplaceAllStringFunc(in, eval)
	fmt.Println(out)
}

func ExampleExercise3() {
	in := strings.Join([]string{
		"다들 그 동안 고생이 많았다.",
		"첫째는 분당에 있는 { 2 ** 4 * 3 }평 아파트를 갖거라.",
		"둘째는 임야 { 10 ** 5 mod 7777 }평을 가져라.",
		"막내는 { 10000 - ( 10 ** 5 mod 7777 ) }평 임야를 갖고",
		"배기량 { 711 * 8 / 9 }cc의 경운기를 갖거라.",
	}, "\n")
	EvalReplaceAll(in)

	// Output:
	// 다들 그 동안 고생이 많았다.
	// 첫째는 분당에 있는 48평 아파트를 갖거라.
	// 둘째는 임야 6676평을 가져라.
	// 막내는 3324평 임야를 갖고
	// 배기량 632cc의 경운기를 갖거라.
}
