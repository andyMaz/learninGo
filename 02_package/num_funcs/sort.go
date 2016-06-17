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

func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		var temp = a[i]
		var j int
		for j = i - 1; j >= 0 && temp < a[j]; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = temp
	}
}