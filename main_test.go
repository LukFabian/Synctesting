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

func TestHandleEcho(t *testing.T) {
	synctest.Run(func() {
		srv, cli := net.Pipe()
		defer srv.Close()
		defer cli.Close()

		go handleEcho(srv)

		body := strings.Repeat("ping\n", 3)

		// Write client data in a goroutine
		go func() {
			cli.Write([]byte(body))
			cli.Close()
		}()

		reader := bufio.NewReader(cli)
		var got strings.Builder

		go func() {
			io.Copy(&got, reader)
		}()

		synctest.Wait()

		if got.String() != body {
			t.Fatalf("echo mismatch: got %q, want %q", got.String(), body)
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
