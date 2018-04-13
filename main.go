package main

import (
	"github.com/tyndyll/pii/adapters"
	"github.com/tyndyll/pii/domain"
	"github.com/tyndyll/pii/processors"
	"github.com/tyndyll/pii/services"
)

func main() {
	validationAdapter := &adapters.ValidationHTTPAdapter{
		ProcessorFunctions: map[string]domain.Processor{
			`ipaddress`:          &processors.IPAddress{},
			`national-insurance`: &processors.NationalInsurance{},
			`name`:               &processors.Name{},
		},
	}
	srv := &services.PiiServer{
		ValidateAdapter: validationAdapter,
	}
	srv.SetupRoutes()
	srv.Start()
}
