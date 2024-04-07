package main

import (
	"log/slog"

	"github.com/tuannkhoi/sport-data-feed/service"
)

func main() {
	logger := slog.Default()

	sdp, err := service.NewSportDataProducer(logger)
	if err != nil {
		logger.Error("Failed to create SportDataProducer: ", err)

		return
	}

	go func() {
		sdp.Monitor()
	}()

	sdp.ProduceNewFootballMatch()
}
