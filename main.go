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
	//var sum, a, b int
	var mas [1000]int
	check_flag:= false
	flag.Parse()
	if *start > *end {
		fmt.Println("Ошибка, start > end")
	} else {
		if *norepeat {
			if *end-*start < *n {
				fmt.Println("Невозможно вывести N чисел от L до R без повторений.")
			} else {
				mas[0] = randInterval(*start, *end)
				for i := 1; i < *n; i++ {
					mas[i] = randInterval(*start, *end)
					for check_flag!= true {
						check_flag= true
						for j := 0; j < i; j++ {
							if mas[i] == mas[j] {
								check_flag= false
							}
						}
						if !check_flag{
							mas[i] = randInterval(*start, *end)
						} else {
							check_flag= true
						}
					}
				}
				for i := 0; i < *n; i++ {
					fmt.Print(mas[i], " ")
				}
			}
		} else {
			for i := 0; i < *n; i++ {
				print(randInterval(*start, *end), " ")
			}
		}
	}
}