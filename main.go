package main

import (
	"bufio"
	"net"
)

func pollChannel(signalChan <-chan struct{}, onSignal func()) {
	go func() {
		<-signalChan
		onSignal()
	}()
}

func handleEcho(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	w := bufio.NewWriter(conn)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		_, err = w.WriteString(line)
		if err != nil {
			return
		}
		w.Flush()
	}
}

func main() {

}
