package main
import "fmt"

// remember if the funtion dont find the value in its 
// local block it will find in the file block 
// so if the variable cant find in the local block 
// it will serch in the global block

var x int = 100



func g(){
	fmt.Println(x);

}

func f(){
	fmt.Println(x);

}

func main(){

	g();
	f();

}
