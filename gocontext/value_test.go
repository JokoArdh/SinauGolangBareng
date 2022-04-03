package gocontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestWithValue(t *testing.T) {
	// parent context
	contextA := context.Background()

	//child
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	//mengambil value di context
	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("b"))
}

//membastalkamn context
func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				//menambahkan simulasi penggunaan context with timeout
				time.Sleep(1 * time.Second)
			}

		}
	}()
	return destination
}

func TestContextCencel(t *testing.T) {
	fmt.Println("Total Gorutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cencel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cencel() //mengirim sinyal cencel ke context setelah prosess selesai

	time.Sleep(2 * time.Second) // harus dikasih time sleep biar proses asysncronoushnya berjalan tidak cepat

	fmt.Println("Total Gorutine : ", runtime.NumGoroutine())
}

//context with time out
func TestContextTimeout(t *testing.T) {
	fmt.Println("Total Gorutine : ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cencel := context.WithTimeout(parent, 5*time.Second)
	defer cencel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	cencel() //mengirim sinyal cencel ke context setelah prosess selesai

	time.Sleep(2 * time.Second) // harus dikasih time sleep biar proses asysncronoushnya berjalan tidak cepat

	fmt.Println("Total Gorutine : ", runtime.NumGoroutine())
}

//context with deadline
func TestContextDeadline(t *testing.T) {
	fmt.Println("Total Gorutine : ", runtime.NumGoroutine())

	parent := context.Background()
	//yang membedakan di sisni digonta ganti sesuai kebutuhan
	ctx, cencel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
	defer cencel()

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
	}

	cencel() //mengirim sinyal cencel ke context setelah prosess selesai

	time.Sleep(2 * time.Second) // harus dikasih time sleep biar proses asysncronoushnya berjalan tidak cepat

	fmt.Println("Total Gorutine : ", runtime.NumGoroutine())
}
