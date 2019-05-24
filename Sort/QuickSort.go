package Sort

func QuickSort(a []int)  {
	len := len(a)
	separateSort(a, 0,len-1)
}

func separateSort(a []int, start,end int)  {
	if start >= end{
		return
	}
	i:= partition(a, start, end)
	separateSort(a,start,i-1)
	separateSort(a, i+1, end)
}

func partition(a []int, start, end int) int {
	// 选取最后一位当对比数字
   	pivot := a[end]

   	var i = start
   	for j := start; j < end; j ++{
   		if a[j] < pivot{
   			if i!=j{
   				a[i], a[j] = a[j],a[i]
			}
			i ++
		}
	}
	a[i],a[end] = a[end],a[i]
	return i
}