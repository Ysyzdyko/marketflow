package collector

import (
	"fmt"
	"log"
	"marketflow/internal/adapters/outbound/exchange/live"
	"marketflow/internal/ports/outbound"
)

var Allowed = map[string]struct{}{
	"BTCUSDT":  {},
	"TONUSDT":  {},
	"DOGEUSDT": {},
	"SOLUSDT":  {},
	"ETHUSDT":  {},
}

func StartAllExchanges() {
	errCh := make(chan error)
	clientCount := 3
	streams := make([]<-chan string, 0, clientCount)

	for i := 1; i <= clientCount; i++ {
		addr := fmt.Sprintf("exchange%d:4010%d", i, i)
		client := live.NewClient(fmt.Sprintf("exchange%d", i), addr, Allowed)
		dataCh := make(chan string, 100)

		go client.Streaming(dataCh, errCh)
		streams = append(streams, dataCh)
	}

	merged := FanIn(streams...)

	wp, err := New(15)
	if err != nil {
		log.Fatal("Worker pool error:", err)
	}
	defer wp.Close()

	for data := range merged {
		dataCopy := data
		err := wp.AddTask(func() {
			rc := outbound.Redis{}
			ProcessData(rc, dataCopy)
		})
		if err != nil {
			log.Println("WorkerPool full:", err)
		}
	}

	go func() {
		for err := range errCh {
			log.Println("Stream error:", err)
		}
	}()
}
