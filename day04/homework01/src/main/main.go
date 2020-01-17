package main

import (
	"fmt"
	"math/rand"
	"time"
)

/* 冒泡排序 */
func bubbleSortInts(a []int, sortOrder func(int,int) bool) {
	length := len(a)
	for top := length - 2; top >= 0; top-- {
		for index := 0; index <= top; index++ {
			if !sortOrder(a[index], a[index+1]) {
				a[index], a[index+1] = a[index+1], a[index]
			}
		}
	}
}

/* 选择排序 */
func selectionSortInts(a []int, sortOrder func(int, int) bool) {
	length := len(a)
	for left := 0; left <= length-2; left++ {
		selectIndex := left
		for index := left + 1; index < length; index++ {
			if !sortOrder(a[selectIndex], a[index]) {
				selectIndex = index
			}
		}
		if selectIndex != left {
			a[selectIndex], a[left] = a[left], a[selectIndex]
		}
 	}
}

/* 插入排序 */
func insertionSortInts(a []int, sortOrder func(int, int) bool) {
	length := len(a)
	for selectIndex := 1; selectIndex < length; selectIndex++ {
		for i := 0; i < selectIndex; i++ {
			if !sortOrder(a[i], a[selectIndex]) {
				tmp := a[selectIndex]
				for j := selectIndex - 1; j >= i; j-- {
					a[j+1] = a[j]
				}
				a[i] = tmp
			}
		}
	}
}

/* 快速排序 */
func quickSortInts(a []int, sortOrder func(int, int) bool) {
	length := len(a)
	if length <= 1 {
		return
	}
	start := 0
	end := length - 1
	mid := a[end]
	left := start
	right := end - 1
	for left < right {
		for sortOrder(a[left], mid) && left < right {
			left++
		}
		for !sortOrder(a[right], mid) && left < right {
			right--
		}
		a[left], a[right] = a[right], a[left]
	}
	if !sortOrder(a[left], a[end]) {
		a[left], a[end] = a[end], a[left]
	}
	quickSortInts(a[:right+1], sortOrder)
	quickSortInts(a[left+1:], sortOrder)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func genSlice() []int {
	arrLen := rand.Intn(10)
	a := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		a[i] = rand.Intn(100)
	}
	return a
}

func testSortOrder (a []int, sortOrder func(int, int) bool) bool {
	arrLen := len(a)
	for i := 1; i < arrLen; i++ {
		if !sortOrder(a[i-1], a[i]) {
			return false
		}
	}
	return true
}

func testBubbleSortInts(testCnt int) {
	sortOrder := func(val1, val2 int) bool { return val1 <= val2}
	for testId := 0; testId < testCnt; testId++ {
		a := genSlice()
		bubbleSortInts(a, sortOrder)
		if !testSortOrder(a, sortOrder) {
			fmt.Println(a)
			panic("bubbleSortInts() test not passed")
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), testId, "bubbleSortInts() passed")
		time.Sleep(time.Millisecond * 10)
	}
}

func testSelectionSortInts(testCnt int) {
	sortOrder := func(val1, val2 int) bool { return val1 <= val2}
	for testId := 0; testId < testCnt; testId++ {
		a := genSlice()
		selectionSortInts(a, sortOrder)
		if !testSortOrder(a, sortOrder) {
			fmt.Println(a)
			panic("selectionSortInts() test not passed")
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), testId, "selectionSortInts() passed")
		time.Sleep(time.Millisecond * 10)
	}
}

func testInsertionSortInts(testCnt int) {
	sortOrder := func(val1, val2 int) bool { return val1 <= val2}
	for testId := 0; testId < testCnt; testId++ {
		a := genSlice()
		insertionSortInts(a, sortOrder)
		if !testSortOrder(a, sortOrder) {
			fmt.Println(a)
			panic("insertionSortInts() test not passed")
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), testId, "insertionSortInts() passed")
		time.Sleep(time.Millisecond * 10)
	}
}

func testQuickSortInts(testCnt int) {
	sortOrder := func(val1, val2 int) bool { return val1 <= val2}
	for testId := 0; testId < testCnt; testId++ {
		a := genSlice()
		quickSortInts(a, sortOrder)
		if !testSortOrder(a, sortOrder) {
			fmt.Println(a)
			panic("quickSortInts() test not passed")
		}
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"), testId, "quickSortInts() passed")
		time.Sleep(time.Millisecond * 10)
	}
}


func main() {
	testBubbleSortInts(1000)
	testSelectionSortInts(1000)
	testInsertionSortInts(1000)
	testQuickSortInts(1000)
}
