package main

import (
	"log/slog"

	"github.com/tuannkhoi/sport-data-feed/service"
)

func main() {
	logger := slog.Default()

	sdc, err := service.NewSportDataConsumer(logger)
	if err != nil {
		logger.Error("Failed to create SportDataConsumer: ", err)

		return
	}

	sdc.Consume()
}
