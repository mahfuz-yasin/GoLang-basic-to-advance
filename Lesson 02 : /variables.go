package main
import "fmt"
func main () {
	var name string = "mahfuz"
	var age int = 32
	// type infarence
	 cgp  := 3.55
	

	fmt.Println("my name is ", name , "i am ", age , "years old", "i got have ", cgp )
	// formatting and use printf
	fmt.Printf("my name is %v i am %v years old i got have %v cgp " ,name ,age,cgp )

	// constant 
	// constant will decleard using capital letter
	const FULLNAME = "mahfuz yasin"
	fmt.Print(FULLNAME)

}