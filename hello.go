/* hello.go - My first Golang program */
package main
import "fmt"
func main() {

    go logMsg("Hello, world\n")
}

func logMsg(msg string){
	fmt.Printf(msg)
}