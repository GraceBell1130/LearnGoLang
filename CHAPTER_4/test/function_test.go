package test

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Example_funcLiteral() {
	func() {
		fmt.Println("Hello!")
	}()

	foo := func() {
		fmt.Println("Literal")
	}
	foo()
	// Output:
	// Hello!
	// Literal
}

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntom\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// ( bill )
	// ( tom )
	// ( jane )
}

func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntom\njane\n")
	var lines []string // 클로저
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
	// Output:
	// [bill tom jane]
}

func NewIntGenerator() func() int {
	var next int
	return func() int {
		next++
		return next
	}
}

func ExampleNewIntGenerator() {
	gen := NewIntGenerator()
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	fmt.Println(gen(), gen(), gen(), gen(), gen())
	gen2 := NewIntGenerator()
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())

	// Output:
	// 1 2 3 4 5
	// 6 7 8 9 10
	// 1 2 3 4 5
}

type NewInt1 int
type NewInt2 int

func foo(i NewInt1) {
	fmt.Println(i)
}

func Example_type() {
	var testInt NewInt1 = 0
	//var testInt2 NewInt2 = 0
	foo(testInt)
	//foo(testInt2) 컴파일 에러
	// Output:
	// 0
}

type BinOp func(int, int) int
type BinSub func(int, int) int

func BinOpToBinSub(f BinOp) BinSub {
	var count int
	return func(a, b int) int {
		fmt.Println(f(a, b))
		count++
		return count
	}
}

func ExampleBinOpToBinSub() {
	sub := BinOpToBinSub(func(a, b int) int {
		return a + b
	})
	/*
		sub := BinOpToBinSub(BinOpToBinSub(func(a, b int) int {
			return a + b
		})) BinSub 타입으로 변경이되므로 컴파일 에러가 발생
	*/
	sub(5, 7)
	sub(5, 7)
	count := sub(5, 7)
	fmt.Println("count:", count)
	// Output:
	// 12
	// 12
	// 12
	// count: 3
}

type StrSet map[string]struct{}
type PrecMap map[string]StrSet

func NewStrSet(strs ...string) StrSet {
	m := StrSet{}
	for _, str := range strs {
		m[str] = struct{}{}
	}
	return m
}

func Eval(opMap map[string]BinOp, prec PrecMap, expr string) int {
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
		if token == "(" {
			ops = append(ops, token)
		} else if _, ok := prec[token]; ok {
			reduce(token)
			ops = append(ops, token)
		} else if token == ")" {
			// 닫는 괄호는 여는 괄호까지 계산을 하고 제거
			reduce(token)
		} else {
			num, _ := strconv.Atoi(token)
			nums = append(nums, num)
		}
	}
	reduce(")") // 초기의 여는 괄호까지 모두 계산
	return nums[0]
}

func NewEvaluator(ooMap map[string]BinOp, prec PrecMap) func(expr string) int {
	return func(expr string) int {
		return Eval(ooMap, prec, expr)
	}
}

func ExampleNewEvaluator() {
	eval := NewEvaluator(map[string]BinOp{
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
	}, PrecMap{
		"**":  NewStrSet(),
		"*":   NewStrSet("**", "*", "/", "mod"),
		"/":   NewStrSet("**", "*", "/", "mod"),
		"mod": NewStrSet("**", "*", "/", "mod"),
		"+":   NewStrSet("**", "*", "/", "mod", "+", "-"),
		"-":   NewStrSet("**", "*", "/", "mod", "+", "-"),
	})

	fmt.Println(eval("5"))
	fmt.Println(eval("1 + 2"))
	fmt.Println(eval("1 - 2 - 4"))
	fmt.Println(eval("( 3 - 2 ** 3 ) * ( -2 )"))
	fmt.Println(eval("3 * ( 3 + 1 * 3 ) / ( -2 )"))
	fmt.Println(eval("3 * ( ( 3 + 1 ) * 3 ) / 2"))
	fmt.Println(eval("1 + 2 ** 10 * 2"))
	fmt.Println(eval("2 ** 3 mod 3"))
	fmt.Println(eval("2 ** 2 ** 3"))

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
}

type VertexID int

func (id VertexID) String() string {
	return fmt.Sprintf("VertexID(%d)", id)
}
func ExampleVertexId_String() {
	i := VertexID(100)
	fmt.Println(i)

	// Output:
	// VertexID(100)
}

type MultiSet map[string]int

func (m MultiSet) Insert(val string) {
	m[val]++
}

func (m MultiSet) Erase(val string) {
	if m[val] <= 1 {
		delete(m, val)
	} else {
		m[val]--
	}
}

func (m MultiSet) Count(val string) int {
	return m[val]
}

func (m MultiSet) String() string {
	s := "{ "
	for val, count := range m {
		s += strings.Repeat(val+" ", count)
	}
	return s + "}"
}

func Example_method() {
	m := MultiSet{}
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	m.Insert("3")
	fmt.Println(m.String())
	fmt.Println(m.Count("3"))
	m.Insert("1")
	m.Insert("2")
	m.Insert("5")
	m.Insert("7")
	m.Erase("3")
	m.Erase("5")
	fmt.Println(m.Count("3"))
	fmt.Println(m.Count("1"))
	fmt.Println(m.Count("2"))
	fmt.Println(m.Count("5"))

	// Output:
	// { }
	// 0
	// { 3 3 3 3 }
	// 4
	// 3
	// 1
	// 1
	// 0
}

type Temp struct {
	X int
}

func (t Temp) GeneralReceiver(i int) {
	t.X += i
}

func (t *Temp) PointerReceiver(i int) {
	t.X += i
}

func Example_Receiver() {
	t := Temp{10}
	t.GeneralReceiver(10) // 20이 아니고 결과값은 10
	// Output:
	// {10}
	t.PointerReceiver(100)
	fmt.Println(t)
	// Output:
	// {110}
}
