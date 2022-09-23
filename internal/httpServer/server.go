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

	serverM.serverHttp = server

	return serverM

}

// processUpdate(context echo.Context) processes Url with metrics data and tries to save it in storage. This is main route function of server.
func (serverM *serverM) processUpdate(context echo.Context) error {
	var retErr error = context.String(http.StatusOK, "Metric saved")
	errCounter := errors.New("wrong type of counter")
	errGauge := errors.New("wrong type of gauge")

	TypeM, NameM, ValueM, err := parseUrl(context)
	if err != nil {
		log.Printf("Got invalid URL in parseUrl(): <%v>", context.Request().URL.Path)
		return context.String(http.StatusBadRequest, fmt.Sprintf("Bad URL: <%v>", context.Request().URL.Path))
	}

	saveErr := serverM.Storage.Save(TypeM, NameM, ValueM)
	if saveErr != nil {
		switch saveErr {
		case errCounter:
			return context.String(http.StatusBadRequest, fmt.Sprintf("Bad counter: <%v>", context.Request().URL.Path))
		case errGauge:
			return context.String(http.StatusBadRequest, fmt.Sprintf("Bad gauge: <%v>", context.Request().URL.Path))
		default:
			return context.String(http.StatusNotImplemented, fmt.Sprintf("Not implemented: <%v>", context.Request().URL.Path))
		}
	}
	return retErr
}

// nonexistentPath(context echo.Context) processes nonexistant paths on server and returns 404 error.
func (serverM *serverM) nonexistentPath(context echo.Context) error {
	log.Printf("Requested nonexistent page: <%v>", context.Request().URL.Path)
	return context.String(http.StatusNotFound, "Page you requested is not found")
}

// parseUrl(context echo.Context) (string, string, string, error) performs parsing of request URL and returns <Type of metric>, <Name of metric>, <Value of metric> and error
func parseUrl(context echo.Context) (string, string, string, error) {
	var url string = context.Request().URL.Path

	urlSlice := strings.Split(url, "/")

	if len(urlSlice) < 5 {
		return "", "", "", errors.New("invalid url")
	}

	return urlSlice[2], urlSlice[3], urlSlice[4], nil
}

// Run() starts server.
func (serverM *serverM) Run() {
	if err := serverM.serverHttp.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
