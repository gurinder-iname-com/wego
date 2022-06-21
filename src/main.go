package main

import (
	"log"
)

func main() {
	log.SetPrefix("WeGo: ")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	IntUnderScoreTest()
	FloatTest()
	// for i := 0; i < 10; i++ {
	// 	log.Println(i)
	// }
}

func IntUnderScoreTest() {
	var firstInt int = 1234567890
	var secondInt = 1_234_567_890
	log.Println(firstInt == secondInt)
}

func FloatTest() {
	var firstFloat float64 = 1.2
	var secondFloat float64 = -1.2
	var thirdFloat float64 = 0
	log.Println(firstFloat)
	log.Println(firstFloat / 0)
	log.Println(secondFloat / 0)
	log.Println(0.0 / thirdFloat)
}
