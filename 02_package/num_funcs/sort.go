package num_funcs

func BubbleSort(a []int) {
	var temp int
	var n = len(a)
	for i := 0; i < n; i++  {
		for j := 1; j < n-i; j++ {
			if a[j-1] > a[j]  {
				temp =  a[j-1]
				a[j-1] = a[j]
				a[j] = temp
			}
		}
	}
}