package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func pollChannel(signalChan <-chan struct{}, onSignal func()) {
	go func() {
		<-signalChan
		onSignal()
	}()
}

func handleEchoAck(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	var sb strings.Builder
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			return
		}
		sb.WriteString(line)
		writer.WriteString(line)
		writer.Flush()
	}
}

func main() {

}
