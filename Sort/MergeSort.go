package Sort

func MergeSort(a []int)  {
	len := len(a)
	if len <=1 {
		return
	}
	mergeSort(a, 0, len-1)
}

func mergeSort(a []int, start, end int)  {
	if start >= end{
		return
	}
	mid := (end + start)/2
	mergeSort(a, start, mid)
	mergeSort(a, mid+1, end)
	merge(a,start,mid,end)

}

func merge(a []int, start, mid, end int)  {
	tmp := make([]int, end-start+1)
	i:= start
	j := mid + 1
	k := 0
	for ; i <= mid && j<= end; k++{
		if a[i] < a[j]{
			tmp[k] = a[i]
			i++
		}else {
			tmp[k] = a[j]
			j++
		}
	}
	for ; i<=mid; i++{
		tmp[k] = a[i]
		k++
	}
	for ; j<=end; j++{
		tmp[k] = a[j]
		k++
	}
	copy(a[start:end+1], tmp)  // 切片不包括最后一位
}

