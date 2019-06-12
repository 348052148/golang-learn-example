package main

import (
	"fmt"
)

func main() {
	s := "我"
	b := "a"
	fmt.Println(len(s))
	fmt.Printf("%b \n", []byte(s))
	//
	//var bytes int
	fmt.Printf("%b \n",0x80) //检查最后1位是否是1
	fmt.Printf("%b \n",s[0]) //检查最后1位是否是1
	fmt.Printf("%b \n", (s[0] << 2));
	fmt.Printf("%b \n", b[0] & 0x80)
	//
	fmt.Println("---")
	fmt.Printf("%b \n", (s[0] << 3))
	fmt.Printf("%b \n", (s[0] << 3) & 0x80)
	fmt.Printf("%b \n", s[0] & 0x80)
	fmt.Println("**************")
	fmt.Printf("%b \n", []byte(s))
	//fmt.Printf("%b \n", (0 | int32((0xe6 & 0x1f) << 16)) | int32((0x88 & 0x7f) << 8) | int32(0x91 & 0x7f))
	//fmt.Printf("%b \n",int32((0x88 & 0x7f) << 8))
	fmt.Printf("%b \n", utf8ToUnicode([]byte(s)))
	fmt.Println("**************")
	fmt.Printf("%b \n",0x88 & 0x3f)

	fmt.Printf("%b \n", 1 << 32)
	fmt.Printf("%b \n", 6 << 1)
}

func utf8ToUnicode(utf8Byte []byte) []byte  {
	var unicodeByte []byte
	var unicode int32 = 0
	var bts uint8
	for _, b := range utf8Byte {
		//4
		if int(b & 0xf0) == 0xf0 {
			unicode |=  int32((b & 0xf) << 18)
			bts = 3
			fmt.Printf("%b :\n", unicode)
			continue
		}
		//3
		if int(b  & 0xe0) == 0xe0 {
			unicode = unicode | int32(b & 0x1f) << 12
			fmt.Printf("%b :\n", unicode)
			bts = 2
			continue
		}
		//2
		if int(b  & 0xc0) == 0xc0 {
			bts = 1
			continue
		}
		unicode = unicode | (int32(b & 0x7f) << (bts * 6))
		bts--
		fmt.Println(bts)
		if bts == 0 {
			fmt.Printf("%b --:\n",unicode)
			var bt []byte
			bt = append(bt,uint8(unicode))
			bt = append(bt,uint8(unicode >> 8))
			bt = append(bt,uint8(unicode >> 16))
			bt = append(bt,uint8(unicode >> 24))
			unicodeByte = append(unicodeByte,bt...)

			unicode = 0
			bts = 0
		}
	}
	return unicodeByte
}