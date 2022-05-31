# 3.1 문자열 

rune형 : int32 정수의 별칭이며 유니코드 포인트 하나를 담을 수 있음
byte형 : uint8의 별칭
nil == empty value ex) nil == nullptr (c++), nil == null(java)

tset 코드 작성 하는 법
func 에 실행할 코드를 작성 후, 예측 결과를 입력한다.
// Output :
// 예측 결과

문자열은 읽기 전용이기 때문에 바이트 조작은 불가능 함
ex) 
s := "가나다"
s[1]++ // error

하지만 문자열을 바이트 단위의 슬라이스로 변환하면 바이트 조작이 가능 함
ex)
b := []byte("가나다") // 문자열을 슬라이스로 형변환 
                      // string(b)는 슬라이스에서 문자열로 형변환
b[1]++ // pass

# 3.2 배열
[...]string{"a","b","c"}처럼 [...]을 이용하여 배열을 생성하게되면 컴파일러가 배열의 요소만큼 할당

n 개의 빈 값을 가지는 슬라이스를 만드는 방법
ex) 
fruits := make([]{변수 타입}, n) // 문자열일 경우 ""으로 초기화, 정수일 경우 0으로 초기화
fruits := make([]{변수 타입}, x, y) // 배열의 길이는 x, 배열의 용량은 y


슬라이싱 : 슬라이스를 자르는 행위 
ex) 
fruits[:1] 

go에서는 python 처럼 음수 값을 이용하여 슬라이싱을 할 수 없음
ex)
fruits[:-1] // error

슬라이스 덧붙이는 방법
fruits = append(fruits, "포도", "딸기") // 요소는 한번에 여러개 덧붙일 수 있음
fruits = append(fruits, other_fruits...) // 슬라이스 배열이 appned 되면 ...를 이용하여 가변 인자를 받는 함수로 호출하여야 함

슬라이스 용량을 확인하는 방법
ex)
x := cap(fruits)

슬라이스 복사하는 방법
ex)
copy(dest, src) // 만약 dest가 src를 복사하기에 용량이 부족하면 용량만큼 copy가 됨

# 3.3 맵
go 언어에서 map은 해시테이블로 구현이 됨
해시맵은 키와 값으로 되어 있으며 키를 이용해서 값을 상수 시간에 가져올 수 있으며 해시맵에는 순서가 없음
해시테이블을 사용하므로 key값이 변경되는 타입을 사용하면 안 됨

맵을 담는 변수를 정의하는 방법
var m map[KeyType]valueType

맵을 생성하는 방법
m := make(map[KeyType]valueType) or m := map[KeyType]valueType{}

값을 읽는 방법
value := m[key] // key가 없을 경우 KeyType이 int일 경우 0, String일 경우 빈 문자열 리턴
value, ok := m[key] // 이와 같이 변수 2개로 리턴 받을 때에는 value에는 위와 같은 값이 ok에는 존재 여부를 bool로 반환 받을 수 있음

값을 쓰는 방법
m[key] = value // key에 값이 있는 경우 변경, 없는 경우 생성

빈 구조체를 값으로 사용하면 값 부분의 메모리를 따로 차지하지 않음
ex) m := m[KeyType]struct{}{} // 값 부분의 메모리를 차지하지 않음

맵에서 데이터 삭제하는 방법
delete(m, key)

# 3.4 입출력
입출력은 io.Reader와 io.Writer 인터페이스와 파생된 다른 인터페이스들을 이용하며 파일뿐만 아니라 버퍼, 소켓 등을 이용하여 읽고 쓸 수 있음

os.Open()은 반환값이 둘이며 하나는 파일 오브젝트 다른 하나는 에러
에러 값이 nil이 되면 파일을 성공적으로 연 것이 됨

defer는 해당 함수를 벗어날 때 호출할 함수를 등록하는 역할 - 함수를 빠져나가는 곳이 한 군데가 아니므로 코드를 좀 더 깔끔하게 작성할 수 있음
fun exam() {
    c, err := os.Create(filename)
    f, err := os.Open(filename)
    defer c.Close() // exam이 종료될 때 실행
    defer f.Close() // exam이 종료될 때 실행
    var num int
    fmt.Fscanf(f, "%d\n", &num)   
    fmt.fprintf(f, "%d\n", num)
} 

