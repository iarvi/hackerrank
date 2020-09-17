#Author

package main
import ( "fmt"
          "strings" )  

func checkstr(str string) {
 estr := []string{}
 ostr := []string{}
     for j:=0; j < len(str); j++{
        if j == 0 || j %2 ==0 {
             estr = append(estr,string(str[j]))
        } else {
         ostr = append(ostr,string(str[j]))
        }
     }
      estr1 := strings.Join(estr,"")
      ostr1 := strings.Join(ostr,"")
      fmt.Println(estr1,ostr1)
      
}
func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT

 var n int
 var str string
 fmt.Scan(&n)
 for i:=0; i< n; i++{
     fmt.Scan(&str)
     checkstr(str)
 }
 }
