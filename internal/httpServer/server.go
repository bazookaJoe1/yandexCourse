package httpserver

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"yandexCourse/internal/storage"

	"log"

	"github.com/labstack/echo"
)

// ServerInit() initializes struct for server and returns it as a pointer.
func ServerInit() *serverM {
	serverM := &serverM{Storage: storage.StorageInit()}

	echoS := echo.New()
	echoS.Any("/*", func(context echo.Context) error {
		log.Printf("ERROR. Got request. URL: %s. Method: %s", context.Request().URL.Path, context.Request().Method)
		return context.String(http.StatusBadRequest, "400 Bad Request: Wrong method")
	})
	echoS.POST("/*", serverM.nonexistentPath)
	echoS.POST("/update/*", serverM.processUpdate)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: echoS,
	}

	serverM.serverHTTP = server

	return serverM

}

// processUpdate(context echo.Context) processes Url with metrics data and tries to save it in storage. This is main route function of server.
func (serverM *serverM) processUpdate(context echo.Context) error {
	errCounter := errors.New("wrong type of counter")
	errGauge := errors.New("wrong type of gauge")

	log.Printf("Got POST request: %s", context.Request().URL.Path)

	TypeM, NameM, ValueM, err := parseURL(context)
	if err != nil {
		log.Printf("Got invalid URL in parseURL(): <%v>", context.Request().URL.Path)
		return context.String(http.StatusNotFound, fmt.Sprintf("Bad URL: <%v>", context.Request().URL.Path))
	}

	saveErr := serverM.Storage.Save(TypeM, NameM, ValueM)
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
	return context.String(http.StatusOK, fmt.Sprintf("Metric <%s> with value <%s> was saved.", NameM, ValueM))
}

// nonexistentPath(context echo.Context) processes nonexistant paths on server and returns 404 error.
func (serverM *serverM) nonexistentPath(context echo.Context) error {
	log.Printf("Requested nonexistent page: <%v>", context.Request().URL.Path)
	return context.String(http.StatusNotFound, "Page you requested is not found")
}

// parseURL(context echo.Context) (string, string, string, error) performs parsing of request URL and returns <Type of metric>, <Name of metric>, <Value of metric> and error
func parseURL(context echo.Context) (string, string, string, error) {
	var url = context.Request().URL.Path

	urlSlice := strings.Split(url, "/")

	if len(urlSlice) < 5 {
		return "", "", "", errors.New("invalid url")
	}

	return urlSlice[2], urlSlice[3], urlSlice[4], nil
}

// Run() starts server.
func (serverM *serverM) Run() {
	if err := serverM.serverHTTP.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
	log.Printf("Server started.")
}
