



// this is variable scope
// it is block wise 
// and inside the block like curly braces
// the value of a variabel in recognizablae
// and outsite the braces the value is not recognizable

package main;

import "fmt"
var x int = 100;
func f(){
	fmt.Printf("%d",x);

}

func g(){
	fmt.Printf("%d",x);
}

func main(){
	f();
	g();
	
}
