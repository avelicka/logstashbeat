package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/radoondas/logstashbeat/beater"
)

func main() {
	err := beat.Run("logstashbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
