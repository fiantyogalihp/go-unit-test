package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

/*
 ! Di dalam go, peraturan menulis func testing, yaitu harus menambahkan kata Test pada awal nama function unit test
*/

type TableTest struct {
	name     string
	request  string
	expected string
}

type BenchmarkTable struct {
	Name    string
	request string
}

/*
 * before & after Test Main
 ! Test Main berjalan disetiap package, jika disalah satu package tidak terdapat TestMain, maka tidak akan dijalankan TestMain nya. Jadi kita harus memberi TestMain pada setiap package ketika TestMain dibutuhkan
*/
func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST DO...")

	m.Run()

	fmt.Println("AFTER UNIT TEST DO...")
}

func BenchmarkTableTest(b *testing.B) {
	benchmarks := []BenchmarkTable{
		{
			Name:    "Fian",
			request: "Fian",
		},
		{
			Name:    "Galih",
			request: "Galih",
		},
		{
			Name:    "TyoPramudya",
			request: "Tyo Pramudya",
		},
		{
			Name:    "Tyo",
			request: "Tyo Pramudya",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

// func BenchmarkSubTest(b *testing.B) {
// 	b.Run("Fian", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("Fian")
// 		}
// 	})
// 	b.Run("Galih", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			HelloWorld("Galih")
// 		}
// 	})
// }

// func BenchmarkHelloWorldFian(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		HelloWorld("Fian")
// 	}
// }
// func BenchmarkHelloWorldGalih(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		HelloWorld("Galih")
// 	}
// }

/*
 ? ketika kita membuat table test, kita tidak perlu membuat ulang unit test yang mayoritas setiap structure nya sama(kombinasi data berbeda-beda), solusinya, kita menggunakan Table Test untuk melakukan testing yang setiap kali membuat unit testing, mayoritas setiap data structure nya sama
*/
func TestTableHelloWorld(t *testing.T) {
	// ? pertama buat data struct table test nya
	tableTest := []TableTest{
		{
			name:     "Fian",
			request:  "Fian",
			expected: "Hello Fian",
		},
		{
			name:     "Tyo",
			request:  "Tyo",
			expected: "Hello Tyo",
		},
		{
			name:     "Galih",
			request:  "Galih",
			expected: "Hello Galih",
		},
		{
			name:     "Pramudya",
			request:  "Pramudya",
			expected: "Hello Pramudya",
		},
	}

	// ? lalu iterasi menggunakan for range untuk melakukan testing setiap data pada struct table test
	for _, test := range tableTest {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

// ? tergantikan oleh table test yang lebih sederhana
// func TestSubTest(t *testing.T) {
// 	t.Run("Eko", func(t *testing.T) {
// 		result := HelloWorld("Eko")
// 		assert.Equal(t, "Hello Eko", result, "Result harus 'Hello Eko'")
// 		fmt.Println("TestSubTest/Eko done")
// 	})
// 	t.Run("Kurniawan", func(t *testing.T) {
// 		result := HelloWorld("Kurniawan")
// 		assert.Equal(t, "Hello Kurniawan", result, "Result harus 'Hello Kurniawan'")
// 		fmt.Println("TestSubTest/Kurniawan done")
// 	})
// }

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("Tidak bisa dijalankan pada Linux")
		// ? func t.Skip() berfungsi membatalkan testing, bukan untuk menggagalkan testing
	}

	result := HelloWorld("Tyo")
	assert.Equal(t, "Hello Tyo", result, "Result harus 'Hello Tyo'") // ! assert akan memanggil t.Fail() ketika gagal
	fmt.Println("TestSkip done")
}

// * testing dengan menggunakan library
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Pramudya")

	require.Equal(t, "Hello Pramudya", result, "Result harus 'Hello Pramudya'") // ! Require akan memanggil t.FailNow() ketika gagal
	fmt.Println("TestHelloWorldRequire done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Fian")

	assert.Equal(t, "Hello Fian", result, "Result harus 'Hello Fian'") // ! assert akan memanggil t.Fail() ketika gagal
	fmt.Println("TestHelloWorldAssert done")
}

// * testing tanpa menggunakan library
// func TestHelloWorldFian(t *testing.T) {
// 	result := HelloWorld("Fian")
// 	if result != "Hello Fian" {
// 		// error
// 		// panic("Result is not a valid 'Hello Fian'") // ! jangan memakai panic untuk mengagalkan unit test, pakai parameter t testing instead
// 		// t.Fail() // ! menggagalkan testing saat itu juga, tetapi kode program dibawahnya akan tetap dijalankan
// 		t.Error("result harus 'Hello Fian'") // ? menampilkan error message lalu menjalankan t.Fail()
// 	}
// 	fmt.Println("TestHelloWorldFian done")
// }

// func TestHelloWorldGalih(t *testing.T) {
// 	result := HelloWorld("Galih")
// 	if result != "Hello Galih" {
// 		// error
// 		// t.FailNow() // ! menggagalkan testing saat itu juga
// 		t.Fatal("result harus 'Hello Galih'") // ? menampilkan error message lalu menjalankan t.FailNow()
// 	}
// 	fmt.Println("TestHelloWorldGalih done")
// }
