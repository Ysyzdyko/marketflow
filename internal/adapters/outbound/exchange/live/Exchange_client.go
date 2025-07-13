package live

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Addr    string
	Name    string
	Filters map[string]struct{}
}

func NewClient(name, addr string, filters map[string]struct{}) *Client {
	return &Client{
		Name:    name,
		Addr:    addr,
		Filters: filters,
	}
}

func (c *Client) Streaming(out chan<- string, errCh chan<- error) {
	conn, err := net.Dial("tcp", c.Addr)
	if err != nil {
		errCh <- fmt.Errorf("failed to connect %s (%s): %w", c.Name, c.Addr, err)
		return
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 3 {
			continue
		}

		symbol := parts[0]
		if _, ok := c.Filters[symbol]; !ok {
			continue
		}

		out <- fmt.Sprintf("%s,%s", c.Name, line)
	}
	if err := scanner.Err(); err != nil {
		errCh <- fmt.Errorf("error reading from %s: %w", c.Name, err)
	}
}
