// it is comparatively hard to 
// find when to deallocate the memory 
// in a certain program lets see a example

package main

import "fmt"

func foo() *int {

	x:=1
	return &x
}

func main (){
	var y *int 
	y = foo()
	fmt.Println(y)
}