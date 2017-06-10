package main

import (
    "fmt"
    "os"
)

func main(){
    var s, sep string
    // use range like for each, the _ is a ignore that return type
    // in rage you return the index and the value, you can not just have
    // it go to an unuesd var as go does not allow
    for _, arg := range os.Args[1:]{
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
