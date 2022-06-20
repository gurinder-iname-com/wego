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
	var firstFloat float32 = 1.2
	log.Println(firstFloat)
}
