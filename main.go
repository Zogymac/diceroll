package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	seed = flag.Int("seed", 0, "seed for random generator. unix(now) be default")
	start = flag.Int("start", 1, "start of range")
	end = flag.Int("end", 6, "end of range")
	n = flag.Int("n", 1, "number of throws")
	norepeat = flag.Bool("norepeat", false, "does it repeat")
)

//Фукнция должна вернуть число из интервала [l,r]
func randInterval(start, end int) int{
	flag.Parse()
	if *seed == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(int64(*seed))
	}
	return rand.Intn(end) + start
}


func main() {
	var sum, a, b int
	flag.Parse()
	if *start > *end {
		fmt.Println("Ошибка, start > end")
	} else {
		if *norepeat {
			if *end - *start < *n {
				fmt.Println("Невозможно вывести N чисел от L до R без повторений.")
			} else {
				a = randInterval(*start, *end)
				sum += a
				fmt.Print(a, " ")
				b = a
				for i := 1; i < *n; i++ {
					a = randInterval(*start, *end)
					sum += a
					sum %= a
					if b != 0 && b != a {
						fmt.Print(a, " ")
						sum += b
					} else {
						i--
					}
					/*
						if sum % a != 0 && sum != b{
							fmt.Println(a, b, sum, sum % a, sum % b)
						} else {
							i--
							sum -= a
							fmt.Println(i, a, b, sum)
						}*/
					b = a
				}
			}
		} else {
			for i := 0; i < *n; i++ {
				print(randInterval(*start, *end), " ")
			}
		}
	}
}