package sighandler

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func SigHandler(exitChan chan int) {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		for {
			s := <-signalChanel
			switch s {
			// kill -SIGHUP XXXX [XXXX - идентификатор процесса для программы]
			//Сигнал SIGHUP отправляется при потере программой своего управляющего терминала.
			case syscall.SIGHUP:
				log.Printf("---> Signal SIGHUP triggered.")

				// kill -SIGINT XXXX или Ctrl+c  [XXXX - идентификатор процесса для программы]
				//Сигнал SIGINT отправляется при введении пользователем в управляющем терминале символа прерывания, по умолчанию это ^C (Control-C).
			case syscall.SIGINT:
				log.Printf("---> Signal SIGINT triggered.")
				exitChan <- 0

				// kill -SIGTERM XXXX [XXXX - идентификатор процесса для программы]
				//SIGTERM — это общий сигнал, используемый для завершения программы.
			case syscall.SIGTERM:
				log.Printf("---> Signal SIGTERM triggered.")
				exitChan <- 0

				// kill -SIGQUIT XXXX [XXXX - идентификатор процесса для программы]
				//Сигнал SIGQUIT отправляется при введении пользователем в управляющем терминале символа выхода, по умолчанию это ^\ (Control-Backslash).
			case syscall.SIGQUIT:
				log.Printf("---> Signal SIGQUIT triggered.")
				exitChan <- 0

			default:
				log.Printf("---> Signal <unknown> triggered.")
				exitChan <- -1
			}
		}
	}()
}
