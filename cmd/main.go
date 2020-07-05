package main

import (
	"github.com/alandtsang/prober/internal/gateway"
	"github.com/labstack/gommon/log"
)

func main() {
	port := 9889

	if err := gateway.New().SetupRoute().ListenAndServe(port); err != nil {
		log.Errorf("prober listen on :%d failed, %v", port, err)
		return
	}
}
