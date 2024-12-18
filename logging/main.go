// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)

// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	log.SetOutput(file)
// 	log.Println("Hello World")

// }

//using logrus Module

package main

import (
	log "github.com/sirupsen/logrus"
)

/*
Trace
Debug
Info
Warn
Error
Fatal
Panic
*/
func main() {
	log.Fatal("Hello World")

}
