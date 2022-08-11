package main

import (
	"context"
	"log"
	"os"
	httpclient "yandexCourse/internal/httpClient"
	metricscollector "yandexCourse/internal/metricsCollector"
	sighandler "yandexCourse/internal/sigHandler"
)

/* initAgent is main function that initializes all business logic*/
func initAgent() {
	var (
		//muGen     sync.Mutex            //mutex for blocking access to *metricscollector.metrics struct
		retChan1  = make(chan struct{}) //index of correct terminating of child contexts
		retChan2  = make(chan struct{})
		exit_chan = make(chan int) //channel for exit code transmission
	)

	ctxPar := context.Background()                        //main background context
	ctxCancMain, cancelMain := context.WithCancel(ctxPar) //main context with cancel

	/*defer func with functional of correct context terminating*/
	defer func() {
		exitCode := <-exit_chan //blocking channel waits for signalHandler to transmit exit code

		log.Printf("---> cancelling main context")

		cancelMain()
		<-retChan1 //blocking channel waits for chlid contexts to terminate
		<-retChan2

		os.Exit(exitCode)
	}()

	metrics := metricscollector.MetricsInit()                                //initialize *metricscollector.metrics struct
	go metricscollector.MetricsCollectorMain(ctxCancMain, metrics, retChan1) //create goroutine for metrics collection

	go sighandler.SigHandler(exit_chan) //create goroutine for signal handling

	agent := httpclient.AgentInit()
	go agent.Run(ctxCancMain, retChan2, metrics)

}

func main() {
	initAgent()
}
