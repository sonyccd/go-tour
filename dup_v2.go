package main

import(
    "bufio"
    "fmt"
    "os"
)

func main(){
    // make a map with strings for keys and ints for values
    //maps are constant time for store and lookup
    counts := make(map[string]int)
    files := os.Args[1:]
    dups := false
    if len(files) == 0 {
        countLines(os.Stdin, counts)
    }else{
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup_v2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        // if the value of that line is great than 1 than print how many duplicates
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
            dups = true
        }
    }
    if !dups{
        fmt.Println("No duplicates")
    }
}

//map is passed as a ref becuase make returns a ref
func countLines(f *os.File, counts map[string]int){
    input := bufio.NewScanner(f)
    for input.Scan(){
        counts[input.Text()]++
    }
}

