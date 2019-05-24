package Sort

// 冒泡排序（从小到大）
func BubbleSort(a []int ){
	len := len(a)
	if len <= 1{
		return
	}

	for i:=0; i<len; i++{
		flag := false
		for j:=0; j<len-i-1; j++{
			if a[j] > a[j+1]{
				a[j], a[j+1] = a[j+1], a[j]
				// 此次冒泡排序有数据交换
				flag =true
			}
		}
		if !flag{
			// 无数据交换，有序
			break
		}
	}
}


