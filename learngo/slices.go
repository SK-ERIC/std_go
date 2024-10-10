package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr length", len(arr))
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])

	arr2 := arr[:3]
	fmt.Println("arr2 length", arr2, len(arr2), arr2[1:])
	arr2[1] = 100
	fmt.Printf("修改了吗, arr1: %d, arr2: %d, \n", arr, arr2)

	// 切片
	/*
		切片是相同类型元素的可变长度的集合，通常表示为[]type。
		每个切片都由三部分组成：
		1. 指向底层数组中某个元素的指针：指向数组的第一个从切片访问的元素，这个元素并不一定是数组的第一个元素。一个底层数组可以对应多个切片，这些切片可以饮用数组的任何位置，并且彼此之间的元素可以重叠
		2. 长度（length/len）：切片中的元素个数
		3. 容量（capacity/cap）：为切片分配的存储空间。
		切片类型的初始化值是nil，没有对应的底层数组，并且长度和容量都为0，所以切片只能和nil比较。如果想检查切片是否为空，则可能使用len(s)==0来判断。
		数组和切片关联的很紧密。切片可以访问数组的全部或着部分元素
	*/
	slice1 := make([]string, 6)
	slice2 := make([]string, 6, 8)
	// 长度都为6，容量分别为6和8
	fmt.Println("slice1", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2", slice2, len(slice2), cap(slice2))

	// 切片的扩容
	foodSlice4 := []string{"肉"}
	fmt.Println(foodSlice4)
	fmt.Println(fmt.Sprintf("len:%d", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "鱼")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "侠")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "鞋")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))
	foodSlice4 = append(foodSlice4, "雨")
	fmt.Println(fmt.Sprintf("Len:%d", len(foodSlice4)))
	fmt.Println(fmt.Sprintf("Cap:%d", cap(foodSlice4)))

	fmt.Println(foodSlice4)
	// 切片的扩容不是改变原切片只想的底层数组，而是生成一个容量更大的底层数组，让后吧原切片中的元素和新元素一起复制到新切片中。
	// 一般来说，可以简单认为新切片的容量是旧切片的2倍。如果原切片的长度大于1024，则新切片的容量是旧切片的1.25倍。
	// 只要心长度不超过原切片的容量，那么使用append函数对其追加元素时，就不会引起扩容。
}
