// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func checkNum(i int) error {
// 	if i%2 == 0 {
// 		return errors.New("Only Odd Numbers Allowed .. ")
// 	}
// 	fmt.Println("Successful Operation")
// 	return nil
// }
// func checkError(e error) {
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// }

// func main() {
// 	e := checkNum(3)
// 	checkError(e)

// }

package main

import (
	"fmt"
)

func checkNum(i int) error {
	if i%2 == 0 {
		return fmt.Errorf("Only Odd Numbers Allowed .. ")
	}
	fmt.Println("Successful Operation")
	return nil
}
func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func main() {
	e := checkNum(3)
	checkError(e)

}
