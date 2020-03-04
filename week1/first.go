// this must be there to run this code

package main


// this is used for formating
import "fmt"


// you can actually alias the variable type
// this is renaming the datatype
// set the type of the dataset
type Celcius float64
type IDnum int

// this is direct 
var x int  = 100
var y = 100
// you can dynamically assign i
// like x:= 100
// but you cant use it outside any function
// you can use it under main  function

func main(){
	// you can use dynamic outsite the function
	fmt.Println("hello world");
	x:=100
	fmt.Println(x);
}
