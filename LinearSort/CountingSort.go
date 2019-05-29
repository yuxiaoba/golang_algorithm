package LinearSort

// 基数排序
func CountingSort(a []int)  {
	num := len(a)
	if num <=1 {
		return
	}
	max := getMax(a)  // 获取待排序数组中的最大值

	c := make([]int, max+1)
	for i:=0; i<num; i++{
		c[a[i]] ++   // 计算每个元素的个数，放入C中
	}

	for i:=1; i<=max; i++{
		c[i] = c[i-1] + c[i]  // 依次累加
	}

	r := make([]int, num)  // 临时数组r，存储排序之后的结果
	for i:=num-1; i>=0; i--{
		index := c[a[i]]-1  // 排序后的位置
		r[index]=a[i]
		c[a[i]]--
	}
	copy(a, r)
}