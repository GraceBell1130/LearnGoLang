package test

import (
	"fmt"
	// "pkg/hangul"
)

func Example_array() {
	fruits := [3]string{"사과", "바나나", "파인애플"}
	for _, fruit := range fruits {
		/*
			if (hangul.HasConsonatSuffix(fruit)) {
				fmt.Printf("%s는 맛있다.\n", fruit)
			}
			else {
				fmt.Printf("%s은 맛있다.\n", fruit)
			}
		*/
		fmt.Printf("%s는 맛있다.\n", fruit)
	}

	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 파인애플는 맛있다.
}

func Example_slicing() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println(nums[1:3])
	fmt.Println(nums[2:])
	fmt.Println(nums[:3])

	// Output:
	// [1 2 3 4 5]
	// [2 3]
	// [3 4 5]
	// [1 2 3]
}

func Example_append() {
	f1 := []string{"사과", "바나나"}
	f2 := []string{"포도", "딸기"}
	f3 := append(f1, f2...)
	f4 := append(f1[:1], f2...)
	fmt.Println(f1)
	fmt.Println(f2)
	fmt.Println(f3)
	fmt.Println(f4)

	// Output:
	// [사과 바나나]
	// [포도 딸기]
	// [사과 바나나 포도 딸기]
	// [사과 포도 딸기]
}

func Example_sliceCap() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(nums)
	fmt.Println("len:", len(nums))
	fmt.Println("cap:", cap(nums))
	fmt.Println()

	sliced1 := nums[:3]
	fmt.Println(sliced1)
	fmt.Println("len:", len(sliced1))
	fmt.Println("cap:", cap(sliced1)) // 뒤에 2개를 잘라내었기 때문에 5
	fmt.Println()

	sliced2 := nums[2:]
	fmt.Println(sliced2)
	fmt.Println("len:", len(sliced2))
	fmt.Println("cap:", cap(sliced2)) // 앞에 2개를 잘라내었기 때문에 3
	fmt.Println()

	sliced3 := nums[:4]
	fmt.Println(sliced3)
	fmt.Println("len:", len(sliced3))
	fmt.Println("cap:", cap(sliced3))
	fmt.Println()

	nums[2] = 100
	fmt.Println(nums, sliced1, sliced2, sliced3) // 동일한 메모리를 바라보고 있기때문에 모든 배열 값이 변경

	//Output:
	// [1 2 3 4 5]
	// len: 5
	// cap: 5
	//
	// [1 2 3]
	// len: 3
	// cap: 5
	//
	// [3 4 5]
	// len: 3
	// cap: 3
	//
	// [1 2 3 4]
	// len: 4
	// cap: 5
	//
	// [1 2 100 4 5] [1 2 100] [100 4 5] [1 2 100 4]
}

func Example_sliceCopy() {
	src := []int{30, 20, 50, 10, 40}
	dest := make([]int, len(src))
	for i := range src {
		dest[i] = src[i]
	}
	fmt.Println(dest)

	// Output:
	// [30 20 50 10 40]
}
