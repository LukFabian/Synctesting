package main

func pollChannel(signalChan <-chan struct{}, onSignal func()) {
	go func() {
		<-signalChan
		onSignal()
	}()
}

func main() {

}
