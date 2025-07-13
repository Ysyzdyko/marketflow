package collector

import (
	"fmt"
	"marketflow/internal/ports/outbound"
	"strings"
	"sync"
)

var (
	batchMu sync.Mutex
	batch   []string
)

// ProcessData обрабатывает строку и батчует запись в Redis
func ProcessData(redis outbound.Redis, d string) {
	batchMu.Lock()

	batch = append(batch, d)
	shouldFlush := len(batch) >= 10
	toWrite := make([]string, len(batch))
	copy(toWrite, batch)

	if shouldFlush {
		batch = nil
	}

	batchMu.Unlock()

	if shouldFlush {
		go flushToRedis(redis, toWrite)
	}
}

func flushToRedis(redis outbound.Redis, data []string) {
	for _, entry := range data {
		parts := strings.Split(entry, ",")
		if len(parts) < 4 {
			fmt.Println("⚠️ invalid data:", entry)
			continue
		}

		symbol := parts[1]
		price := parts[2]

		if err := redis.Set(symbol, price); err != nil {
			fmt.Println("❌ failed to write to Redis:", err)
		}
	}
}
func FanIn(channels ...<-chan string) <-chan string {
	out := make(chan string)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(c <-chan string) {
			defer wg.Done()
			for data := range c {
				out <- data
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
