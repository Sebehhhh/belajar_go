package main

import "fmt"

func main() {
    // Variabel dan Konstanta
    var a int = 10
    var b int = 5
    const nilaiMax int = 100

    // Operator
    var c int

    // Operator aritmatika
    fmt.Println("Penambahan:", a+b)          // Output: Penambahan: 15
    fmt.Println("Pengurangan:", a-b)         // Output: Pengurangan: 5
    fmt.Println("Perkalian:", a*b)           // Output: Perkalian: 50
    fmt.Println("Pembagian:", a/b)           // Output: Pembagian: 2
    fmt.Println("Sisa bagi:", a%b)           // Output: Sisa bagi: 0

    // Operator penugasan
    c = a + b
    fmt.Println("Nilai c setelah ditugaskan:", c) // Output: Nilai c setelah ditugaskan: 15

    // Operator perbandingan
    fmt.Println("a == b:", a == b)           // Output: a == b: false
    fmt.Println("a != b:", a != b)           // Output: a != b: true
    fmt.Println("a > b:", a > b)             // Output: a > b: true
    fmt.Println("a < b:", a < b)             // Output: a < b: false
    fmt.Println("a >= b:", a >= b)           // Output: a >= b: true
    fmt.Println("a <= b:", a <= b)           // Output: a <= b: false

    // Operator logika
    var x bool = true
    var y bool = false
    fmt.Println("x && y:", x && y)          // Output: x && y: false
    fmt.Println("x || y:", x || y)          // Output: x || y: true
    fmt.Println("!x:", !x)                  // Output: !x: false

    // Variabel Konstan
    fmt.Println("Nilai maksimum:", nilaiMax) // Output: Nilai maksimum: 100
}
