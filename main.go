package main

import (
	"fmt"
	"strings"
)

const (
	FIRST  uint32 = 4278190080
	SECOND uint32 = 16711680
	THIRD  uint32 = 65280
	FORTH  uint32 = 255

	DEFAULT = "0.0.0.0/0"
)

var rtable map[uint32]uint32 = map[uint32]uint32{
	ip(10, 1, 2, 0):     ip(10, 1, 0, 2),
	ip(10, 1, 1, 0):     ip(10, 1, 1, 1),
	ip(172, 20, 0, 0):   ip(172, 20, 0, 2),
	ip(172, 20, 100, 0): ip(172, 20, 100, 4),
}

func main() {
	var ip uint32 = ip(172, 20, 100, 52)
	fmt.Println("next router is", FindNextRouter(ip))
}

func FindNextRouter(ip uint32) string {
	maxLen := 0
	var nextRouter uint32

	for k, v := range rtable {
		bitLen := strings.Count(bit(k&ip), "1")
		if maxLen < bitLen {
			maxLen = bitLen
			nextRouter = v
		}
	}
	return ipv4(nextRouter)
}

func ip(a, b, c, d uint32) uint32 {
	return a<<24 + b<<16 + c<<8 + d
}

func ipv4(ip uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d",
		FIRST&ip>>24, SECOND&ip>>16, THIRD&ip>>8, FORTH&ip)
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

func getNetwork(ip, smask uint32) uint32 {
	return ip & smask
}
