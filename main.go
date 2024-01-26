package main

import "fmt"

type ByteSize int

const (
    _           = iota 				// ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota) 	// left shifting == 2^10 == 1024
    MB 								// MB ByteSize = 2 << (10 * iota)
    GB
    TB
    PB
    EB
)

func main() {
	fmt.Printf("%d \t\t\t %b\n", KB, KB)
	fmt.Printf("%d \t\t %b\n", MB, MB)
	fmt.Printf("%d \t\t %b\n", GB, GB)
	fmt.Printf("%d \t\t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
}	