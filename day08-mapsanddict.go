#Author : iArvi

package main
import ("fmt"
        "os"
        "bufio")

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
var n, phno int
var name string
var sstr[] string
scanner := bufio.NewScanner(os.Stdin)
phonebook := make(map[string]int)
    fmt.Scan(&n)
    for i:=0; i< n; i++{
        fmt.Scan(&name, &phno)
        phonebook[name] = phno     
    }
//its good idea to split taking input and checking with map.
    for scanner.Scan(){
        sstr=append(sstr,scanner.Text())
    }
    for i:= range(sstr){
        if value, ok := phonebook[sstr[i]]; ok {
            fmt.Printf("%s=%d\n", sstr[i], value)
        } else {
            fmt.Println("Not found")
            
        }
    }
} 
