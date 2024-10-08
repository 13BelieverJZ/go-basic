package main

import "fmt"

func main() {
	var arr1 [5]int
	fmt.Println(arr1)

	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)

	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr4)

	arr_5 := [...]int{0: 3, 1: 5, 4: 6}
	fmt.Println(arr_5)

	//二维数组
	var arr_6 = [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_6)

	arr_7 := [3][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {3, 4, 5, 6, 7}}
	fmt.Println(arr_7)

	arr_8 := [...][5]int{{1, 2, 3, 4, 5}, {9, 8, 7, 6, 5}, {0: 3, 1: 5, 4: 6}}
	fmt.Println(arr_8)

	// note :数组不可动态变化问题，一旦声明了，其长度就是固定的
	var arr = [5]int{1, 2, 3, 4, 5}
	//arr[5] = 6
	fmt.Println(arr)

	// 1.数组是值类型问题，在函数中传递的时候是传递的值，如果传递数组很大，这对内存是很大开销。
	modifyArr(arr)
	fmt.Println(arr) //[1 2 3 4 5]

	// 参数数组地址
	modifyArr2(&arr)
	fmt.Println(arr) //[1 100 3 4 5]

	// 2.数组赋值问题，同样类型的数组（长度一样且每个元素类型也一样）才可以相互赋值，反之不可以
	//var arrX = [5]int{1, 2, 3, 4, 5}
	//var arr_1 [5]int = arrX
	//var arr_2 [6]int = arrX
	//fmt.Println(arr_1, arr_2, arrX)
}

func modifyArr(arr [5]int) {
	arr[1] = 100
	fmt.Println("modifyArr:", arr)
}

func modifyArr2(arr *[5]int) {
	arr[1] = 100
	fmt.Println("modifyArr2:", arr)
}
