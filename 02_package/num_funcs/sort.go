package num_funcs

func BubbleSort(a []int) {
	var n = len(a)
	for i := 0; i < n; i++  {
		for j := 1; j < n-i; j++ {
			if a[j-1] > a[j]  {
				a[j-1], a[j] = a[j], a[j-1] //simple swap
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

// easier just to perform simple swap
func swap(x, y int) (int, int) {
	return y, x
}