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

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func main(){
	fmt.Printf("%08b\n", FlagUp)
	fmt.Printf("%08b\n", FlagBroadcast)
	fmt.Printf("%08b\n", FlagLoopback)
	fmt.Printf("%08b\n", FlagPointToPoint)
	fmt.Printf("%08b\n", FlagMulticast)
	
	fmt.Println(YiB) //how do I print this?
}
