package main

import "fmt"

func main() {
	// Membuat slice dengan 5 data awal
	data := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice awal:", data)

	// Menambahkan 3 data ke dalam slice
	data = append(data, 6, 7, 8)

	fmt.Println("Slice setelah penambahan data:", data)
}
