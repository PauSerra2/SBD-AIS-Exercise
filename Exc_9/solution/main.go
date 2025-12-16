package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"

    "exc9/mapred"
)

func main() {
    // Read meditation.txt
    f, err := os.Open("res/meditations.txt")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Load lines in []string as chunks 
    var chunks []string
    sc := bufio.NewScanner(f)
    // augments the buffer in case 
    buf := make([]byte, 0, 64*1024)
    sc.Buffer(buf, 1<<20)
    for sc.Scan() {
        chunks = append(chunks, sc.Text())
    }
    if err := sc.Err(); err != nil {
        panic(err)
    }

    // Executes MapReduce
    var mr mapred.MapReduce
    results := mr.Run(chunks)

    // Prints the top 50 words by frequency
    type kv struct {
        K string
        V int
    }
    list := make([]kv, 0, len(results))
    for k, v := range results {
        list = append(list, kv{K: k, V: v})
    }
    sort.Slice(list, func(i, j int) bool {
        if list[i].V == list[j].V {
            return list[i].K < list[j].K
        }
        return list[i].V > list[j].V
    })

    top := 50
    if len(list) < top {
        top = len(list)
    }
    for i := 0; i < top; i++ {
        fmt.Printf("%s %d\n", list[i].K, list[i].V)
    }
}
