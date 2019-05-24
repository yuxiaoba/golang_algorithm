package Sort

// 插入排序（从小到大）
func InsertSort(a []int)  {
	len := len(a)

	if len<=1{
		return
	}

	for i:=1; i<len; i++{
		val := a[i]
		j := i-1
		for ; j>=0; j--{
			if a[j] > val{
				a[j+1] = a[j]
			}else {
				break
			}
		}
		// a[j] <= val，在j后面插入
		a[j+1] = val


	}
}
