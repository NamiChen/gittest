package hello

import (
	"fmt"
	"testing"
)

const Pi float64 = 3.1415
const Monday, Tuesday, Wednesday, Thursday, Friday, Saturday = 1, 2, 3, 4, 5, 6
const a = iota
const (
	b = iota
	c
)

var aaa = 15

var aa, bb, cc = 1, 2, 3

var (
	name = "ckai"
	bbb  = false
)

//复数
var c1 complex64 = 5 + 19i

const (
	Monday1, Tuesday2, Wednesday3 = 1, 2, 3
	Thursday4, Friday5, Saturday6 = 4, 5, 6
)

func TestHello(t *testing.T) {
	aaa = 2
	want := "Hello, world."
	if got := hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
func Test() {
	fmt.Println(a)
}
