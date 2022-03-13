package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//function untuk membuat sebuah TEST
func TestHello(t *testing.T) {
	result := Helloword("joko")

	if result != "Hello joko" {
		//eror
		//panic("hasil dari variabel salah")
		//eror menggagalkan unit test
		//t.Fail()

		t.Error("Harusnya joko") //fungsi eror

	}
	fmt.Println("Test hellow word done !")

}

func TestHellojoko(t *testing.T) {
	result := Helloword("jack")

	if result != "Hello jack" {
		//eror
		//panic("hasil dari variabel salah")
		//menggagalkan unit test
		//t.FailNow()

		t.Fatal("result must 'hello jack' ")
	}

	fmt.Println("Test hellow word joko !")
}

//TESTING MENGGUNAKAN TESTIFY LIBRARY GITHUB
func TestAssertion(t *testing.T) {
	result := Helloword("joko")

	assert.Equal(t, "Hello joko", result, "hasilnya kudune helo joko lo !") //baris parameter terakhir adalah message

	fmt.Println("TEST hello word dengan assertion selesai dieksekusi")
}

// menggunakan parameter Testify Reuire
func TestAssertionRequire(t *testing.T) {
	result := Helloword("joko")

	require.Equal(t, "Hello joko", result, "test menggunakan require hallo joko")

	fmt.Println("test dengan require selesai")
}

//skip unit test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("tidak dapat berjalan di Windows")
	}

	//testingnya
	result := Helloword("joko")
	require.Equal(t, "Hello joko", result, "test menggunakan require hallo joko")

}

//menggunakan before dan after unit testing
func TestMain(m *testing.M) {
	//before unit test
	fmt.Println("BEFORE unit test dijalankan")

	//mengeksekusi semua unit test
	m.Run()

	//after unit test
	fmt.Println("AFTER unit test dijalankan")
}

//menggunakan sub test
func TestSubtes(t *testing.T) {
	t.Run("joko", func(t *testing.T) {
		result := Helloword("joko")
		require.Equal(t, "Hello joko", result, "test menggunakan require hallo joko")
	})

	t.Run("ardi", func(t *testing.T) {
		result := Helloword("ardi")
		require.Equal(t, "Hello ardi", result, "test menggunakan require hallo ardi")
	})
}

//menggunakan tabel tes
func TestHelloTable(t *testing.T) {
	tes := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "joko",
			request:  "joko",
			expected: "Hello joko",
		},
		{
			name:     "ketket",
			request:  "ketket",
			expected: "Hello ketket",
		}, {
			name:     "aldy",
			request:  "aldy",
			expected: "Hello aldy",
		},
	}

	for _, test := range tes {
		t.Run(test.name, func(t *testing.T) {
			result := Helloword(test.request)
			require.Equal(t, test.expected, result)
		})

	}
}

//mengiji kecepatann dan performa kode dengan benchmark
// func BenchmarkHello(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Helloword("joko")
// 	}
// }
// func BenchmarkHelloName(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Helloword("ardiyanto")
// 	}
// }

//menggunakan sub benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("joko", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Helloword("joko")
		}
	})
	b.Run("ardiyanto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Helloword("ardiyanto")
		}
	})
}
