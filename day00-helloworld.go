#Author : iarvi

package main
import ("fmt"
        "bufio" 
        "os")

func main() {
fmt.Println("Hello, World.")
reader := bufio.NewReader(os.Stdin)
inputString,_ := reader.ReadString('\n')
fmt.Println(inputString)
}
