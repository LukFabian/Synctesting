package main

import (
	"context"
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
