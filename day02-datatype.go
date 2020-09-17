#Author: iarvi
package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
 
    // Declare second integer, double, and String variables.
    var x uint64
    var f float64
    var s1 string
    // Read and save an integer, double, and String to your variables.
   // reader := bufio.NewReader(os.Stdin)
    fmt.Scan(&x)
    fmt.Scan(&f)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        s1 = scanner.Text()
    }
    // Print the sum of both integer variables on a new line.
    fmt.Println(i+x)
    
    // Print the sum of the double variables on a new line.
    fmt.Printf("%.1f",d + f )
    fmt.Printf("\n")
    
    // Concatenate and print the String variables on a new line
    // The 's' variable above should be printed first.
    fmt.Println(s + s1)
    
}
