package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"yandexCourse/internal/storage"

	"log"

	"github.com/labstack/echo"
)

// ServerInit() initializes struct for server and returns it as a pointer.
func ServerInit() *ServerM {
	ServerM := &ServerM{Storage: storage.StorageInit()}

	echoS := echo.New()
	echoS.Any("/*", func(context echo.Context) error {
		log.Printf("ERROR. Got request. URL: %s. Method: %s", context.Request().URL.Path, context.Request().Method)
		return context.String(http.StatusMethodNotAllowed, "405 Method Not Allowed")
	})
	echoS.POST("/*", ServerM.nonexistentPath)
	echoS.POST("/update/*", ServerM.processUpdate)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: echoS,
	}

	ServerM.serverHTTP = server

	return ServerM

}

// processUpdate(context echo.Context) processes Url with metrics data and tries to save it in storage. This is main route function of server.
func (ServerM *ServerM) processUpdate(context echo.Context) error {
	typeM := context.Param("type")
	nameM := context.Param("name")
	valueM := context.Param("value")

	errCounter := errors.New("wrong type of counter")
	errGauge := errors.New("wrong type of gauge")

	log.Printf("Got POST request: %s", context.Request().URL.Path)

	saveErr := ServerM.Storage.Save(typeM, nameM, valueM)
	if saveErr != nil {
		switch saveErr.Error() {
		case errCounter.Error():
			return context.String(http.StatusBadRequest, fmt.Sprintf("Bad counter: <%v>", context.Request().URL.Path))
		case errGauge.Error():
			return context.String(http.StatusBadRequest, fmt.Sprintf("Bad gauge: <%v>", context.Request().URL.Path))
		default:
			return context.String(http.StatusNotImplemented, fmt.Sprintf("Not implemented: <%v>", context.Request().URL.Path))
		}
	}
	return context.String(http.StatusOK, fmt.Sprintf("Metric <%s> with value <%s> was saved.", nameM, valueM))
}

// nonexistentPath(context echo.Context) processes nonexistant paths on server and returns 404 error.
func (ServerM *ServerM) nonexistentPath(context echo.Context) error {
	log.Printf("Requested nonexistent page: <%v>", context.Request().URL.Path)
	return context.String(http.StatusNotFound, "Page you requested is not found")
}

// Run() starts server.
func (ServerM *ServerM) Run() {
	if err := ServerM.serverHTTP.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Printf("Server started.")
}
