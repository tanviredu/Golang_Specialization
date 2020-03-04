

// pointer is something that shows the address of the memory
// a pointer is an address to a memory
// it mostly  the virtuyal memory
// two things is very important
// '&' reference the address
// '*' dereference the memory
// you put the & before the variable it return the memory address
// and if you put the * before the memory address it will show the value

package main

import "fmt"

func main(){
	var x int = 1
	var y int
	var ip *int // ip isa  pointer
	ip  = &x // ip points to the memory address of the value of the 
	y = *ip
	fmt.Println(x);
	fmt.Println(ip);
	fmt.Println(y);

}
