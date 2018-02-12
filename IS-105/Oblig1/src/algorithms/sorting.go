package main

//importere filen her
import (
	"fmt"
)


// Implementering av Bubble_sort modified algoritmen
func Bubble_sort_modified(a []int) {
	var (
		n = len(a)
		sorted = false
	)
	for !sorted {
		swapped := false
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				swapped = true
			}
		}
		if !swapped {
			sorted = true
		}
		n = n - 1
	}
}


// Implementering av Bubble_sort algoritmen
func BubbleSort(list []int) {
	// find the length of list n
	n := len(list)
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if list[j] > list[j+1] {
				temp := list[j+1]
				list[j+1] = list[j]
				list[j] = temp
			}
		}
	}
}

// Implementering av Quicksort algoritmen
func QuickSort(values []int) {
	qsort(values, 0, len(values)-1)
}

func qsort(values []int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot > values[j] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	qsort(values, l, i-2)
	qsort(values, i, r)
}


// Her er Main funksjonen til å printe array lister
func main() {
	list := [] int {1,3,2,4,8,6,7,2,3,0} // Arrays will fall for Bubble Sort Function
	fmt.Println("Unsorted array: " , list)
	BubbleSort(list)
	fmt.Println("Sorted array: ", list, "\n",)

	values := [] int {1,4,2,3,7,6,8,5,0} // This Array list will affiliate with Quick Sort function
	fmt.Println("Values before Quick sorting: ", values)
	QuickSort(values)
	fmt.Println("Values after Quick sorting: ", values, "\n",)

	a := []int{5, 1, 6, 2, 4, 8, 3, 7, 9} //Arrays named 'a' affiliate with QuickSort modified function
	fmt.Println("QuickSort_Modified Array Before Sorting : ", a)
	Bubble_sort_modified(a)
	fmt.Println("Sorted Array : ", a, "\n",)

}
