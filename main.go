package main

import (
    "os"
    "fmt"
)

func main(){
    for _, e := range os.Environ() {
        fmt.Println(e)
    }

}
