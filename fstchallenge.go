package main

import "fmt"

const ebulicaoK float64 = 373

func main() {

	var tempK = ebulicaoK
	var tempC = tempK - 273 // conversao de K para C

	fmt.Printf("A temperatura de ebulição da água em K=%v e em C= %v.", tempK, tempC)
}
