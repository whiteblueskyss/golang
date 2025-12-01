package main

import (
	"fmt"
	"sort"
)

// slice -> dynamic
// most used construct in go
// + useful methods

func slice() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)

	//fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...) // this 3 dots means to unpack the second slice and append its values. extracting everything appending to the slice. So second index is not included.

	fmt.Println(courses)

}

// uninitialized slice is nil

// var nums = make([]int, 0, 5)
// capacity -> maximum numbers of elements can fit
// fmt.Println(cap(nums))
// fmt.Println(nums == nil)

// nums = append(nums, 1)
// nums = append(nums, 2)

// slice operator

// var nums = []int{1, 2, 3, 4, 5}
// fmt.Println(nums[0:1])
// fmt.Println(nums[:2])
// fmt.Println(nums[1:])

// slices
// var nums1 = []int{1, 2, 3}
// var nums2 = []int{1, 2, 4}

// fmt.Println(slices.Equal(nums1, nums2))

// var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
// fmt.Println(nums)
