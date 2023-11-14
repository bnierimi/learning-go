package main

import (
	"fmt"
	// "math/rand"
	"time"
	"sync"
)

var mutex = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id0", "id1", "id2", "id3", "id4"}

var allResults = []string{}

func main(){
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\n(+) Results: %v", allResults)
}

func dbCall(id int) {
	// var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[id])
	log()
	wg.Done()
}

func save(result string) {
	mutex.Lock()
	allResults = append(allResults, result)
	mutex.Unlock()
}

func log() {
	mutex.RLock()
	fmt.Printf("\nThe current results are: %v", allResults)
	mutex.RUnlock()
}