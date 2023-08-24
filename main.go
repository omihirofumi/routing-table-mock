package main

import "fmt"

func main() {
	ipA := ip(255, 255, 255, 0)
	fmt.Println(bit(ipA))
}

func ip(a, b, c, d uint32) uint32 {
	return a<<24 + b<<16 + c<<8 + d
}

func bit(d uint32) string {
	var res string
	for d > 0 {
		if d%2 == 1 {
			res = "1" + res
		} else {
			res = "0" + res
		}
		d /= 2
	}
	return res
}
