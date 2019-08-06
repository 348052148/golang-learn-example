package main

import "fmt"

func main() {

	var a uint32 = 41
	////100
	fmt.Printf("%b \n", a)
	fmt.Printf("%b \n", a>>1)
	fmt.Printf("%b \n", a&1)

	fmt.Printf("%b \n", a>>1>>1)
	fmt.Printf("%b \n", (a>>1)&1)
	var b uint32=0
	fmt.Printf("a %b \n", b|1)
	fmt.Printf("a %b \n", (b|1) << 1)
	fmt.Printf("a %b \n", (((b+1) << 1)|1) << 1)

	//var n uint32 = 0
	//fmt.Printf("%b \n", n << 1)
	//fmt.Printf("%b \n", a & 1)
	//n +=a & 1
	//a >>= 1
	//n +=a & 1
	//a >>= 1
	//n +=a & 1
	//a >>= 1
	//n +=a & 1
	//fmt.Printf("%b \n", a & 1)
	//fmt.Printf("%b \n", n)

	fmt.Printf("%b", reverseBits(41))
}
//00000010100101000001111010011100
func reverseBits(num uint32) uint32 {
	times := 32
	var res uint32 = 0
	for times > 0 {
		res<<= 1
		res += num & 1
		//res += num % 2
		num>>=1
		times--
	}
	return res
}