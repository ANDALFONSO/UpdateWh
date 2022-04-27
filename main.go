package main

import (
	"UPDATE_WH/handlers"
	"UPDATE_WH/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/fury_go-core/pkg/transport/httpclient"
)

type Dependencies struct {
	updateSercice service.IUpdateService
}

func main() {
	server := gin.Default()
	b := build()
	handlers.UpdateHandler(server, b.updateSercice)
	server.Run(os.Getenv("PORT"))
}

func build() *Dependencies {

	wdmTimeout, _ := time.ParseDuration("2m")
	wdmRetryMax := 1
	optsWdm := []httpclient.OptionRetryable{
		httpclient.WithTimeout(wdmTimeout),
		httpclient.EnableCache(),
	}
	hTTPClient := httpclient.NewRetryable(wdmRetryMax, optsWdm...)

	updateService := service.NewPS(hTTPClient)
	return &Dependencies{
		updateSercice: updateService,
	}

}
