package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}
func Sqrt(x float64) (float64, float64, int) {
	// var z, z_prev float64 = 1.0, 0.0
	var z, z_prev float64 = x / 2, 0.0
	times := 1
	for {
		z -= (z*z - x) / (2*z)
		if math.Abs(z - z_prev) < 0.01 {
			break
		}
		z_prev = z
		times += 1
	}
	return z, (z - math.Sqrt(x)), times
}

func main() {
	defer fmt.Println("world")

	sum := 1
	
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum)

	// for ; sum < 1000; {
	// 	sum += sum
	// }

	// for sum < 1000 {
	// 	sum += sum
	// }
	for {
		sum += sum
		if sum >= 1000 {
			break
		}
	}
	fmt.Println(sum)

	fmt.Println(sqrt(2), sqrt(-4))

	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	fmt.Println(Sqrt(1.0))
	fmt.Println(Sqrt(2.0))
	fmt.Println(Sqrt(3.0))

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println(time.Saturday)
	fmt.Println(time.Now().Weekday())

	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("GMo")
	case t.Hour() < 17:
		fmt.Println("GAf")
	default:
		fmt.Println("GEv")
	}

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Hello")
}