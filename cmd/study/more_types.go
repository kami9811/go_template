package main

import (
	"fmt"
	"strings"
	"math"
	// "golang.org/x/tour/pic"
	// "golang.org/x/tour/wc"
)

type Vertex struct {
	// X int
	// Y int
	X, Y int
}
type VertexFloat struct {
	Lat, Long float64
}
// var m map[string]VertexFloat = map[string]VertexFloat{
// 	"Bell Labs": VertexFloat{
// 		40.6, -74.3,
// 	},
// 	"Google": VertexFloat{
// 		37.4, -122.0,
// 	},
// }
var m map[string]VertexFloat = map[string]VertexFloat{
	"Bell Labs": {
		40.6, -74.3,
	},
	"Google": {
		37.4, -122.0,
	},
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

var pow_slice []int = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Pic(dx, dy int) [][]uint8 {
	var frame [][]uint8
	for i := 0; i < dy; i++ {
		row := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			row[j] = uint8((i + 1) * (j + 1))
		}
		frame = append(frame, row)
	}
	return frame
}
func WordCount(s string) map[string]int {
	var m map[string]int = make(map[string]int)
	array := strings.Split(s, " ")
	for _, value := range array {
		_, exist := m[value]
		if exist {
			m[value] += 1
		} else {
			m[value] = 1
		}
	}
	return m
}
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
func fibonacci() func() int {
	primary_value := 0
	secondary_value := 1
	temporary_value :=0
	return func() int {
		primary_value = temporary_value
		temporary_value = primary_value + secondary_value
		secondary_value = primary_value
		return primary_value
	}
}


func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	// fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
	p0 := &v
	p0.X = 1e9
	fmt.Println(v)
	
	fmt.Println(v1, v2, v3, p)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var s []int = primes[1:4]
	primes[1] = 1e4
	fmt.Println(s)

	q := []int{2, 3, 5, 11, 13}
	fmt.Println(q)
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	st := []struct {
		i int
		b bool
	}{
		{2, true},
		{13, true},
	}
	fmt.Println(st)

	var s0 []int
	fmt.Println(s0, len(s0), cap(s0))
	if s0 == nil {
		fmt.Println("nil!")
	}

	sa := make([]int, 5)
	printSlice("sa", sa)
	sb := make([]int, 0, 5)
	printSlice("sb", sb)
	sc := sb[:2]
	printSlice("sc", sc)
	sd := sc[2:5]
	printSlice("sd", sd)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	var sappend []int
	printSlice("sappend", sappend)
	sappend = append(sappend, 0)
	printSlice("sappend", sappend)
	sappend = append(sappend, 1)
	printSlice("sappend", sappend)
	sappend = append(sappend, 2, 3, 4, 5, 6, 7, 8)
	printSlice("sappend", sappend)

	// for i, v := range pow_slice {
	// 	fmt.Printf("2**%d = %d\n", i, v)
	// }
	for _, value := range pow_slice {
		fmt.Printf("%d\n", value)
	}
	for index := range pow_slice {
		fmt.Printf("index: %d\n", index)
	}

	var frame [][]uint8 = Pic(3, 3)
	fmt.Println(frame)
	// pic.Show(Pic)

	// m = make(map[string]VertexFloat)
	// m["Bell Labs"] = VertexFloat{
	// 	40.6, -74.3,
	// }
	// fmt.Println(m["Bell Labs"])
	fmt.Println(m)

	mutationg := make(map[string]int)

	mutationg["Answer"] = 42
	fmt.Println("The value: ", mutationg["Answer"])
	mutationg["Answer"] = 48
	fmt.Println("The value: ", mutationg["Answer"])
	delete(mutationg, "Answer")
	fmt.Println("The value: ", mutationg["Answer"])
	answer_value, answer_ok := mutationg["Answer"]
	fmt.Println("The value: ", answer_value, "Present?", answer_ok)

	// wc.Test(WordCount)

	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))

	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println("fibonacci: ", fib())
	}
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x,
	)
}