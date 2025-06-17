package main

import (
	"bufio"
	"context"
	"io"
	"net"
	"strings"
	"testing"
	"testing/synctest"
)

func TestAfterFunc(t *testing.T) {
	synctest.Run(func() {
		ctx, cancel := context.WithCancel(context.Background())

		funcCalled := false
		context.AfterFunc(ctx, func() {
			funcCalled = true
		})

		synctest.Wait()
		if funcCalled {
			t.Fatalf("AfterFunc function called before context is canceled")
		}

		cancel()

		synctest.Wait()
		if !funcCalled {
			t.Fatalf("AfterFunc function not called after context is canceled")
		}
	})
}

func TestHandleEchoAck(t *testing.T) {
	synctest.Run(func() {
		srv, cli := net.Pipe()
		defer srv.Close()
		defer cli.Close()

		go handleEchoAck(srv)

		body := strings.Repeat("ping\n", 3)

		acked := make(chan struct{})

		go func() {
			cli.Write([]byte(body))
			cli.(net.Conn).Close()
			buf := bufio.NewReader(cli)
			io.ReadAll(buf)
			close(acked)
		}()

		synctest.Wait()

		select {
		case <-acked:
			// Good: client finished ACK
		default:
			t.Fatalf("ACK handler did not complete")
		}
	})
}

func TestPollChannel(t *testing.T) {
	synctest.Run(func() {
		signalChan := make(chan struct{})

		called := false
		pollChannel(signalChan, func() {
			called = true
		})

		synctest.Wait()
		if called {
			t.Fatalf("pollChannel called func before signal")
		}

		// Send the signal
		signalChan <- struct{}{}

		synctest.Wait()
		if !called {
			t.Fatalf("pollChannel did not call func after signal")
		}
	})
}
