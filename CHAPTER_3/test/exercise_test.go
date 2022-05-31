package test

import "fmt"

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonatSuffix(s string) bool {
	numEnds := 28
	result := false
	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index&numEnds != 0
		}
	}

	return result
}

func Example_Exercise_array() {
	fruits := [3]string{"사과", "바나나", "메론"}
	for _, fruit := range fruits {
		if HasConsonatSuffix(fruit) {
			fmt.Printf("%s는 맛있다.\n", fruit)
		} else {
			fmt.Printf("%s은 맛있다.\n", fruit)
		}
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 메론은 맛있다.
}

func quick_sort(target []int, left int, right int) {
	pivot := target[(left+right)/2]
	row := left
	hight := right
	for row < hight {
		for target[row] < pivot {
			row++
		}
		for pivot < target[hight] {
			hight--
		}
		if row <= hight {
			target[row], target[hight] = target[hight], target[row]
			row++
			hight--
		}
	}
	if left < hight {
		quick_sort(target, left, hight)
	}
	if row < right {
		quick_sort(target, row, right)
	}
}

func Example_Exercise_sort() {
	target := []int{3, 4, 1, 2, 5}
	quick_sort(target, 0, len(target)-1)
	fmt.Println(target)

	// Output:
	// [1 2 3 4 5]
}

func binary_search(target []string, value string, start int, last int) bool {
	middle := (start + last) / 2
	if target[middle] == value {
		return true
	} else if middle == start || middle == last {
		return false
	}
	if target[middle] < value {
		return binary_search(target, value, middle+1, last)
	}
	if target[middle] > value {
		return binary_search(target, value, start, middle-1)
	}
	return false
}

func Example_Exercise_Check_string() {
	target := []string{"2", "9", "11", "15", "28", "33", "40", "47", "51", "64", "76", "94"}
	fmt.Println(binary_search(target, "51", 0, len(target)-1))
	fmt.Println(binary_search(target, "80", 0, len(target)-1))

	// Output:
	// true
	// false
}
func pop(queue *[]int) int {
	if 0 < len(*queue) {
		var pop_data int = (*queue)[0]
		*queue = (*queue)[1:]
		return pop_data
	}
	return -1
}

func push(queue *[]int, value int) {
	*queue = append(*queue, value)
}

func Example_Exercise_Queue() {
	var queue []int
	push(&queue, 4)
	push(&queue, 5)
	push(&queue, 7)
	fmt.Println(pop(&queue))
	fmt.Println(pop(&queue))
	fmt.Println(pop(&queue))
	fmt.Println(pop(&queue))

	// Output:
	// 4
	// 5
	// 7
	// -1
}

func NewMultiSet() map[string]int {
	return make(map[string]int)
}

func Insert(m map[string]int, val string) {
	map_value, checker := m[val]
	if checker {
		m[val] = map_value + 1
	} else {
		m[val] = 1
	}
}

func Erase(m map[string]int, val string) {
	map_value, checker := m[val]
	if checker {
		if map_value-1 <= 0 {
			delete(m, val)
		} else {
			m[val] = map_value - 1
		}
	}
}

func Count(m map[string]int, val string) int {
	return_value, checker := m[val]
	if checker {
		return return_value
	}
	return 0
}

func String(m map[string]int) string {
	return_string := "{"
	for keys := range m {
		for i := 0; i < m[keys]; i++ {
			return_string += (" " + keys)
		}
	}
	return_string += "}"
	return return_string
}

func Example_Multi_Set() {
	m := NewMultiSet()
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	Insert(m, "3")
	fmt.Println(String(m))
	fmt.Println(Count(m, "3"))
	Insert(m, "1")
	Insert(m, "2")
	Insert(m, "5")
	Insert(m, "7")
	Erase(m, "3")
	Erase(m, "5")
	fmt.Println(Count(m, "3"))
	fmt.Println(Count(m, "1"))
	fmt.Println(Count(m, "2"))
	fmt.Println(Count(m, "5"))

	// Output:
	// {}
	// 0
	// { 3 3 3 3}
	// 4
	// 3
	// 1
	// 1
	// 0
}
