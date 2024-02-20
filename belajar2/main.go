package main

import "fmt"

func main() {
	// Deklarasi variabel dengan tipe data string
	var nama string = "John Doe"

	// Deklarasi variabel dengan tipe data int
	var umur int = 30

	// Menampilkan nilai variabel
	fmt.Println("Nama:", nama)
	fmt.Println("Umur:", umur)

	fmt.Println("=============================================")
	// Deklarasi variabel dengan tipe data string
	var name string

	// Penugasan ulang variabel name dengan nilai baru yang memiliki tipe data string
	name = "John Doe"

	// Deklarasi variabel dengan tipe data int
	var age int

	// Penugasan ulang variabel age dengan nilai baru yang memiliki tipe data int
	age = 30

	// Menampilkan nilai variabel name dan age
	fmt.Println("Nama:", name)
	fmt.Println("Umur:", age)

	// Penugasan ulang variabel name dan age dengan nilai baru yang memiliki tipe data yang sama
	name = "Jane Smith"
	age = 35

	// Menampilkan nilai variabel name dan age setelah ditugaskan kembali
	fmt.Println("Nama Baru:", name)
	fmt.Println("Umur Baru:", age)

	fmt.Println("=============================================")
	// Inferensi tipe untuk variabel saldo
	saldo := 1000.50

	// Inferensi tipe untuk variabel aktif
	aktif := true

	// Menampilkan tipe data dari variabel saldo dan aktif
	fmt.Printf("Tipe data dari variabel saldo: %T\n", saldo)
	fmt.Printf("Tipe data dari variabel aktif: %T\n", aktif)

	fmt.Println("=============================================")
	// Inferensi tipe untuk beberapa variabel
	a, b := 10, "Hello"

	// Menampilkan tipe data dari variabel a dan b
	fmt.Printf("Tipe data dari variabel a: %T\n", a)
	fmt.Printf("Tipe data dari variabel b: %T\n", b)

	fmt.Println("=============================================")
	// Variabel underscore (_) digunakan untuk menangani nilai yang tidak diperlukan
	_, status := getPersonInfo()

	// Menampilkan status saja, sedangkan nama tidak digunakan
	fmt.Println("Status:", status)

	fmt.Println("=============================================")
	// Deklarasi variabel dengan nilai yang akan diformat
	namaMahasiswa := "John Doe"
	umurMahasiswa := 20
	jurusan := "Informatika"
	ipk := 3.75

	// Menggunakan printf dengan nama lain untuk memformat dan menampilkan nilai variabel
	fmt.Printf("Nama Mahasiswa: %s\n", namaMahasiswa)
	fmt.Printf("Umur Mahasiswa: %d\n", umurMahasiswa)
	fmt.Printf("Jurusan: %s\n", jurusan)
	fmt.Printf("IPK: %.2f\n", ipk)

}
func getPersonInfo() (string, string) {
	return "John Doe", "Aktif"
}
