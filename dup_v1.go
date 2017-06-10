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
    // buffered input from std input
    input := bufio.NewScanner(os.Stdin)
    // scan one line at a time
    //returns ture if there is input, false if none
    fmt.Println("Ready")
    for input.Scan() {
        //will add a map entry for that line and count up if already there
        //this works because ints are set to 0 on init
        counts[input.Text()]++
    }
    for line, n := range counts {
        // if the value of that line is great than 1 than print how many duplicates
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }else{
            fmt.Println("No duplicates")
        }
    }
}

