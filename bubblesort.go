package main

import "fmt"

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				// Tukar posisi elemen
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	// Contoh array yang akan diurutkan
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	fmt.Println("Array sebelum diurutkan:", arr)

	bubbleSort(arr)

	fmt.Println("Array setelah diurutkan:", arr)
}
