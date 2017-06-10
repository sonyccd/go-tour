//read in the file and look for duplicates
//read the entire file inot mem then search

package main

import (
    "fmt"
    "os"
    "strings"
    "io/ioutil"
)

func main(){
    counts := make(map[string]int)
    for _, filename := range os.Args[1:]{
        data, err := ioutil.ReadFile(filename)
        if err != nil{
            fmt.Fprintf(os.Stderr ,"dup_v3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n"){
            counts[line]++
        }
    }
    for line, n := range counts{
        if n > 1{
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}
