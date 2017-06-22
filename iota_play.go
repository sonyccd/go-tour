package main

import "fmt"

type Flags uint

const(
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main(){
	fmt.Printf("%08b\n", FlagUp)
	fmt.Printf("%08b\n", FlagBroadcast)
	fmt.Printf("%08b\n", FlagLoopback)
	fmt.Printf("%08b\n", FlagPointToPoint)
	fmt.Printf("%08b\n", FlagMulticast)
}
