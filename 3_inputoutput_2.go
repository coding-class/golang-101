package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Text 1 -> ")
    text, _ := reader.ReadString('\n')        
    fmt.Println(text)

    fmt.Print("Text 2 -> ")
    text2, _ := reader.ReadString('\n')        
    fmt.Println(text2)
}
