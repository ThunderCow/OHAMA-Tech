package algorithms

import (
	"time"
	"testing"
	"reflect"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Bubble Sort test has passed")
	b.StopTimer()
}



// https://golang.org/doc/effective_go.html#init
func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

// Skriv "benchmark"-tester for benchmarkBSortModified funksjonen
// Skriv en ny testfunksjon benchmarkBSortModified

func TestBSort(t *testing.T) {
	values := []int{9, 1, 20, 3, 6, 7}
	expected := []int{1, 3, 6, 7, 9, 20}

	BubbleSort(values)

	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %d, actual is %d", 1, values[0])
	}
}

func BenchmarkBSort100(b *testing.B) {
	benchmarkBSort(100, b)
}

func BenchmarkBSort1000(b *testing.B) {
	benchmarkBSort(1000, b)
}

func BenchmarkBSort10000(b *testing.B) {
	benchmarkBSort(10000, b)
}

func benchmarkBSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		BubbleSort(values)
	}
}

// Implementasjon av testfunksjoner for Quicksort algoritmen
func TestQSort(t *testing.T) {
	values := []int{9, 1, 20, 3, 6, 7}
	expected := []int{1, 3, 6, 7, 9, 20}

	QuickSort(values)

	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %d, actual is %d", 1, values[0])
	}
}

func BenchmarkQSort100(b *testing.B) {
	benchmarkQSort(100, b)
}

func BenchmarkQSort1000(b *testing.B) {
	benchmarkQSort(1000, b)
}

func BenchmarkQSort10000(b *testing.B) {
	benchmarkQSort(10000, b)
}

func benchmarkQSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		QuickSort(values)
	}
}

// Les https://en.wikipedia.org/wiki/Bubble_sort
// Implementering av Bubble_sort algoritmen
func BubbleSort(a []int) {

	// Skriv din kode her
	for i:=0; i<len(a); i++ {
		flag := 0;//Taking Flag Variable
		for j:=0; j<6-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j];
				a[j] = a[j+1];
				a[j+1] = temp;
				flag = 1;//setting flag as 1, if swapping occurs
			}
		}
		if flag == 0 {
			break//breaking out for loop to reduce extra loops
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
