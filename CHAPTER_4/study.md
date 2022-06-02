# 4.1 값 남겨주고 넘겨받기

슬라이스는 Reference Type(Slice, Map, String)이며 구조가 배열에 대한 포인터, 길이, 용량 이렇게 세 값으로 이루어져있다.

만약 아래와 같이 슬라이스를 매개변수로 하는 함수를 호출하면 동일한 배열에 대한 포인터, 길이, 용량을 가진 슬라이스가 생성가 되어 호출한 함수 내부에서 값을 변경 후

리턴이 되어도 슬라이스의 배열 값이 변경된 상태로 유지가 된다.

**주의 할 점이 원본이 넘어가는 것이 아니라 동일한 배열에 대한 포인터 값과 길이 용량을 가진 슬라이스가 복사되어 생성된 다는 것을 주의해야 한다.**

```
func foo1(para int[]) // 일반적으로 매개변수 넘기기 <- Reference Type이 아닌 함수를 넘기면 Call by Value 규약을 따름
func foo2(para *int[]) // 포인터를 활용하여 매개변수 넘기기 <- Call by Referece 규약을 따름
```

위 같은 방법으로 작성 된 함수에서 값을 변경하게 되면 변수 원본이 매개변수로 넘어간다.

&을 이용하면 변수의 주소값을 알 수 있으며 *을 앞에 붙이면 값을 참조 할 수 있음

함수의 반환 값을 둘 이상으로 만들고 싶으면 괄호로 둘러싸고 ,로 구분하여 반환 타입을 작성해주면 된다.

```
func foo() (int, string) {

}

func caller() {
   int_value, string_value = foo() // 값을 받을 때에도 ,로 구분하여 반환값에 수에 맞게 받으면 됨
}
```

Go에서는 돌려주는 값들 역시 넘겨 받는 인자와 같은 형태로 쓸 수 있으며 해당 인자들은 기본값으로 초기화가 됨

```
func foo() (n int, s string) { // n은 0으로 초기화, s는 빈 문자열로 초기화
    n = 1
    s = "test"
    retrun // 1과 test가 리턴이 됨
}
```

가변인자로 데이터를 넘기는 법
```
func foo(lines... string) { 

}

func caller() {
    foo("Hello", "World") // 슬라이스 Type으로 가변인자가 생성 됨
    lines := []string{"Go Language", "c++ Languate"}
    foo(lines...) // 슬라이스 Type을 넘길 때에는 ...을 사용하여 넘겨야 함
}
```

# 4.2 값으로 취급되는 함수

리터럴 함수 선언 방법 // c++ 람다랑 비슷한 개념

```
func (a int, b int) int {
    return a + b
}

func Example() {
    add := func (a int, b int) int {
        return a + b
    }

    add(1, 3)

    func (a int, b int) int {
        return a - b
    } (3, 2) // 리터럴 함수를 변수에 담지 않고 직접 호출할때에는 리터럴 함수에 ()를 붙여서 인자를 넘겨줘야한다.
}
```

고계 함수(higer-order function) 

```
func foo(output string ,f func(s string)) { // 이 처럼 함수를 매개변수로 하는 함수를 고계 함수라고 함
    f(output)
}

func Example() {
    s := "Hello"
    foo(func(s string){ // 리터럴 함수를 넣어서 고계 함수인 foo 함수 호출
        fmt.Println(s)
    })
}
```

클로저 (closure) - 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드를 의미

```
func foo(output string, f func(s string)) {
    f(output)
}
fucn Example() { 
    var lines []string
    output := "Hello"
    foo(output, func(s string){
        lines = append(lines, s)   // 리터럴 함수 외부에 있는 lines 변수를 함수 내부에서 사용을 하고 있음
    })

}
```

생성기

```
func NewIntGenerator() func() int {
	var next int
	return func() int { // 고계 함수를 반환
		next++
		return next
	}
}

func foo () {
    gen := NewIntGenerator() 
	fmt.Println(gen(), gen(), gen(), gen(), gen())
    gen2 := NewIntGenerator() // 같은 생성기를 이용하여 만들지만 접근하는 클로저 값(next)은 다름 
	fmt.Println(gen2(), gen2(), gen2(), gen2(), gen2())
}
```

Named Type 만드는 방법은 type 예약어를 사용하여 만들 수 있다.

명명된 자료형을 만들면 컴파일 시간에서 동일한 자료형을 쓰는 함수를 호출할 때, 잘 못 호출하는 것을 방지할 수 있다.

변수 타입 뿐만아니라 함수도 같은 효과를 볼 수 있다.

```
type NewInt1 int
type NewInt2 int
type runes rune[]
func foo(i NewInt1) int{
    ...
}

func caller() {
    var testInt NewInt2 = 0
    foo (testInt) // 컴파일 에러
    rune[] = runes{"1", "2"} // 이런 식으로 명명된 자료형과 명명되지 않는 자료형은 호환이 가능
}
```
# 4.3 메서드

메서드 : 리시버 매개변수를 가지는 함수
```
func (recv T) MethodName(p1 T1, p2 T2) R1
        └ 리시버 부분
```

포인터 리시버 : 자료형이 포인트형인 리시버
일단 리시버일 경우에는 메서드 내부의 값이 변경되더라도 호출자에게 변경된 값은 반영되지 않지만 포인터 리시버일 경우 메서드 내부의 값이 변경되면 호출자에게 변경된 값들이 반영 된다.
```
type Temp struct {
    X int
}
func (t Temp) GeneralRecvier(i int) {
    t.X += i
}

func (t Temp) PointerRecvier(i int) {
    t.X += i
}

func caller() {
    t := Temp(10)
    t.GeneralRecvier(10)
    t.GeneralRecvier(100)
    fmt.Println(t)
    // Output:
    // 110
}
```

메서드의 이름이 대문자로 시작하면 해당 메서드는 다른 모듈에게 보이기 때문에 호출이 가능
메서드의 이름이 소문자로 시작되면 다른 모듈에서는 보이지 않음
메서드가 아닌 함수, 명명된 자료형도 마찬가지로 위와 같음
)
