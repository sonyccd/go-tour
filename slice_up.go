package main

import "fmt"

func main() {
	months := [...]string{1: "Jan", 2: "Feb", 3: "March", 4: "Apr", 5: "May", 6: "Jun", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println("Q2:", Q2)
	fmt.Println("summer:", summer)
}
