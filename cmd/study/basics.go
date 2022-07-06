package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

// basic types
var (
	ToBe   bool	      = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

// variables
// var c, python, java bool
// var i, j int = 1, 2

// const
const Pi = 3.14
const (
	Big = 1 << 100
	Small = Big >> 99
)

func main() {
	// random
	fmt.Println(rand.Seed)
	fmt.Println(rand.Intn(10))

	// math
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi)

	// function
	fmt.Println(add(42, 13))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(17))

	// variables
	// var i int
	// var c, python, java = true, false, "no!"
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)

	// basic types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// type conversion
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	// var f float64 = math.Sqrt(x*x + y*y)
	var z uint = uint(f)
	// var z uint = f
	fmt.Println(x, y, z)

	// type inference
	v := 3.14 + 0.5i
	fmt.Printf("v is of type %T\n", v)

	// const
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	
}