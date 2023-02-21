package main

import "fmt"

func main() {
	var str string
	str = "Go 爱好者"
	//The string："Go 爱好者"
	fmt.Printf("The string：%q\n", str)
	// runes(char)： ['G' 'o' ' ' '爱' '好' '者']
	fmt.Printf("runes(char)： %q\n", []rune(str))
	// runes(char)： ['G' 'o' ' ' '爱' '好' '者']
	fmt.Printf("runes(char)： %q\n", []rune(str))
	//runes(byte)： [47 6f 20 e7 88 b1 e5 a5 bd e8 80 85]
	fmt.Printf("runes(byte)： [% x]\n", []byte(str))
	for byte_index, char_val := range str {
		fmt.Printf("%d : %q [% x]\n", byte_index, char_val, []byte(string(char_val)))
	}

	for c2 := range str {
		fmt.Println(c2)
	}
	/*
		0 : 'G' [47]
		1 : 'o' [6f]
		2 : ' ' [20]
		3 : '爱' [e7 88 b1]
		6 : '好' [e5 a5 bd]
		9 : '者' [e8 80 85]
	*/
}
