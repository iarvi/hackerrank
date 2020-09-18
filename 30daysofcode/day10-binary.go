#Author : iArvi

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func main() {
    var count, maxcount int
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int64(nTemp)
   // bin := strconv.FormatInt(n, 2)
    count = 0
    for n > 0 {
         r := n % 2
            n= n/2
            if(r ==1){
                count+=1
                if(count>maxcount){
                    maxcount = count
                }
            }else{
                count = 0;
            }
    }
    //fmt.Println(bin)
    fmt.Println(maxcount)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

