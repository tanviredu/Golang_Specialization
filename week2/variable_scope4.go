// variable scope is the area
// where the variable is accesabe
// you will see an error in the code 
// because the value x is inside the function f()

// so when the g() search for it it will not found in its block
// but then it will search in the file block and still dont find it
// so it will return an error

package main

import "fmt"


func f(){
	var x int = 4;
	fmt.Println(x);
}


func g(){

	fmt.Println(x);
}

func main(){

	f();
	g();
}
