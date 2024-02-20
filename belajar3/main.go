package main

import (
	"fmt"
)

func main() {
	// Variabel dengan nilai 3.63
	var nilai1 float32 = 3.63

	// Variabel dengan nilai 10.12345
	var nilai2 float32 = 10.12345

	// Variabel dengan nilai 7.5
	var nilai3 float32 = 7.5

	// Menggunakan fmt.Printf untuk memformat dan mencetak nilai-nilai float dengan berbagai variasi angka di belakang koma
	fmt.Printf("Nilai float 1: %.2f\n", nilai1) // Menampilkan dua angka di belakang koma
	fmt.Printf("Nilai float 2: %.3f\n", nilai2) // Menampilkan tiga angka di belakang koma
	fmt.Printf("Nilai float 3: %.4f\n", nilai3) // Menampilkan empat angka di belakang koma
	fmt.Println("=============================================")

	// Deklarasi variabel bool dengan nilai true
	var isTrue bool = true
	fmt.Printf("Variabel isTrue: %t\n", isTrue)

	// Deklarasi variabel bool dengan nilai false
	var isFalse bool // Variabel ini akan memiliki nilai default (zero value) yaitu false
	fmt.Printf("Variabel isFalse: %t\n", isFalse)

	// Nilai zero (default) dari tipe data bool
	var defaultBool bool
	fmt.Printf("Nilai default bool: %t\n", defaultBool)
	fmt.Println("=============================================")

	// Pointer
	var ptr *int
	fmt.Println("Nilai pointer:", ptr) // Output: Nilai pointer: <nil>

	// Slice
	var s []int
	fmt.Println("Nilai slice:", s) // Output: Nilai slice: []

	// Map
	var m map[string]int
	fmt.Println("Nilai map:", m) // Output: Nilai map: map[]

	// Channel
	var ch chan int
	fmt.Println("Nilai channel:", ch) // Output: Nilai channel: <nil>

	// Interface kosong
	var emptyInterface interface{}
	fmt.Println("Nilai interface kosong:", emptyInterface) // Output: Nilai interface kosong: <nil>
}
