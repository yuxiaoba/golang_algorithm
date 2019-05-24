package Sort

func SeletcSort(a []int)  {
	len := len(a)
    if len <=1 {
    	return
	}
	for i:=0; i<len; i++{
		min := i
		for j:=i+1; j<len; j++{
			if a[j] < a[min]{
				min = j
				}
		}

		a[i],a[min] = a[min],a[i]

	}
}