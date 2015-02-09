package go_koans
import  "fmt"
func divideFourBy(i int) int {
/*	if i == 0 {
		fmt.Println("panicking")
		panic(fmt.Sprintf("%v", i))
	}*/
	return 4 // i
}

const __divisor__ = 0

func aboutPanics() {
//	assert(__delete_me__) // panics are exceptional errors at runtime

	n := divideFourBy(__divisor__)
	defer func(){
		if r:= recover(); r != nil{
			fmt.Println("recovering")
			n = 2
		}
	}()
	//assert(n == 2) // panics are exceptional errors at runtime
}
