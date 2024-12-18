package process

import "fmt"

func CheckEven(i int) string {
	if i%2 == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println(CheckEven(2))
}
