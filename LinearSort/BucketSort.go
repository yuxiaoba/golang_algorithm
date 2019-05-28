package LinearSort

import (
	"fmt"
	"algorithm/Sort"
)

// 桶排序

// 获取待排序数组中的最大值
func getMax(a []int)int{
	max := a[0]
	for i:=1; i<len(a); i++{
		if a[i] > max{
			max = a[i]
		}
	}
	return max
}


func BucketSort(a []int)  {
	num := len(a)
	max := getMax(a)
	buckets := make([][]int, num)  // 二维切片

	index :=0
	for i:=0; i< num; i++{
		index = a[i]*(num -1) / max  // 桶序号
		buckets[index] = append(buckets[index],a[i]) // 加入对应的桶中
	}

	tmpPos := 0  // 标记数组位置
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0{
			Sort.QuickSort(buckets[i])  // 桶内做快速排序
			copy(a[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	fmt.Println(a)
}


// 桶排序简单实现
func BucketSort1(source []int)  {
	var array [20]int
	for i:=0; i<len(source); i++{
		array[source[i]] ++
	}

	for i:=0; i<len(array); i++{
		for array[i] != 0 {
			fmt.Printf("%v ", i)
			array[i] --
		}
	}
	fmt.Println()
}

