package mapred

import (
    "regexp"
    "strings"
    "sync"
)

type MapReduce struct{}

// Run executes the pipeline Map → Shuffle → Reduce in a concurrent way
func (mr MapReduce) Run(input []string) map[string]int {
    // 1) MAP in parallel per chunk
    mappedChan := make(chan []KeyValue, len(input))
    var wgMap sync.WaitGroup
    wgMap.Add(len(input))

    for _, chunk := range input {
        c := chunk
        go func() {
            defer wgMap.Done()
            mapped := mr.wordCountMapper(c)
            mappedChan <- mapped
        }()
    }

    go func() {
        wgMap.Wait()
        close(mappedChan)
    }()

    // 2) SHUFFLE: take all the KeyValue and agrupate for key []int
    grouped := make(map[string][]int)
    for batch := range mappedChan {
        for _, kv := range batch {
            grouped[kv.Key] = append(grouped[kv.Key], kv.Value)
        }
    }

    // 3) REDUCE: add values for key in parallel
    result := make(map[string]int, len(grouped))
    var mu sync.Mutex
    var wgReduce sync.WaitGroup
    // limit concurrency(fix now)
    sem := make(chan struct{}, 8)

    for k, vals := range grouped {
        key := k
        values := vals

        wgReduce.Add(1)
        sem <- struct{}{}
        go func() {
            defer wgReduce.Done()
            kv := mr.wordCountReducer(key, values)
            mu.Lock()
            result[kv.Key] = kv.Value
            mu.Unlock()
            <-sem
        }()
    }

    wgReduce.Wait()
    return result
}

// wordCountMapper: take the words in alfabetical, lowercase and assigns value 1
func (mr MapReduce) wordCountMapper(text string) []KeyValue {
    // només lletres (filtra numerals i especials)
    reWord := regexp.MustCompile(`[A-Za-z]+`)
    normalized := strings.ToLower(text)
    tokens := reWord.FindAllString(normalized, -1)

    out := make([]KeyValue, 0, len(tokens))
    for _, w := range tokens {
        out = append(out, KeyValue{Key: w, Value: 1})
    }
    return out
}

// wordCountReducer: adds all the values for the key
func (mr MapReduce) wordCountReducer(key string, values []int) KeyValue {
    sum := 0
    for _, v := range values {
        sum += v
    }
    return KeyValue{Key: key, Value: sum}
}
