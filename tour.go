package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

// Oh this is wonderful, look at this declaration
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

type box struct {
	width  int
	height int
	//volume
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	x := 20
	ar := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	fmt.Printf("x is of %T\n", x)
	sum := func(a, b int) int { return a + b }(3, 4)
	fmt.Println(sum)
	first := "test"
	second := "testing"
	fmt.Println(swap(first, second))
	fmt.Println(upDown(100))
	fmt.Println(summ(ar))
	fmt.Println(pow(3, 2, 10))
	box_test := make_box(2, 3)
	fmt.Printf("%T\n", box_test)
	fmt.Println(box_test.height, box_test.width)
	fmt.Println(squrt(10.0))
	fmt.Println(math.Sqrt(10))
}

// look mom I can return more than one thing :p
func swap(a string, b string) (string, string) {
	return b, a
}

// this is a naked return
func upDown(sum int) (x, y int) {
	x = sum - 1
	y = sum + 1
	return
}

func summ(a []int) int {
	var out int
	for i := 0; i < len(a); i++ {
		out += a[i]
	}
	return out
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func readword(ch chan string) {
	fmt.Println("Type a word, then hit Enter.")
	var word string
	fmt.Scanf("%s", &word)
	ch <- word
}

func timeout(t chan bool) {
	time.Sleep(5 * time.Second)
	t <- true
}

//make a box and return a pointer to it
// go will just put it on the heap for you
func make_box(w, h int) *box {
	temp := box{w, h}
	return &temp
}

func squrt(v float64) float64 {
	approx := 0.5 * v
	var delta float64 = 1
	for delta > 0.0001 {
		temp_approx := 0.5 * (approx + v/approx)
		delta = math.Abs(temp_approx - approx)
		approx = temp_approx
	}
	return approx
}

func cuncur() {
	t := make(chan bool)
	go timeout(t)

	ch := make(chan string)
	go readword(ch)

	select {
	case word := <-ch:
		fmt.Println("Received", word)
	case <-t:
		fmt.Println("Timeout.")
	}
}
