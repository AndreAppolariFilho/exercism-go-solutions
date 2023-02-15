package letter

import (
	"runtime"
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func AuxiliaryFrequency(m FreqMap, s string) {
	for _, r := range s {
		m[r]++
	}
}
func ConcurrentFrequency(l []string) FreqMap {
	wordStream := make(chan string)
	go func() {
		for _, s := range l {
			wordStream <- s
		}
		close(wordStream)
	}()

	var wg sync.WaitGroup
	freqMapStream := make(chan FreqMap)

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			for s := range wordStream {
				freqMapStream <- Frequency(s)
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(freqMapStream)
	}()

	result := make(FreqMap)

	for fm := range freqMapStream {
		for r, c := range fm {
			result[r] += c
		}
	}

	return result
}
