package heavystruct

import (
	"fmt"
	"testing"
)

func writeToChannel(ch chan HeavyStruct, data HeavyStruct) {
	ch <- data
}

func writePointerToChannel(ch chan *HeavyStruct, data *HeavyStruct) {
	ch <- data
}

func BenchmarkWriteToChannel(b *testing.B) {
	data := createHeavyStructWithRandomValues()
	ch := make(chan HeavyStruct)
	// Create a goroutine to consume data from the channel
	go func() {
		for receivedData := range ch {
			// Consume data from the channel by printing it
			fmt.Println(receivedData)
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writeToChannel(ch, data)
	}
}

func BenchmarkWritePointerToChannel(b *testing.B) {
	data := createHeavyStructWithRandomValues()
	ch := make(chan *HeavyStruct)
	// Create a goroutine to consume data from the channel
	go func() {
		for receivedData := range ch {
			fmt.Println(*receivedData)
		}
	}()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		writePointerToChannel(ch, &data)
	}
}
