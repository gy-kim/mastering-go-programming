package hydrarestapi

import (
	"log"
	"net/http"

	"github.com/gy-kim/mastering-go-programming/hydra/hydraconfigurator"
)

type DBlayerconfig struct {
	DB   string `json:"database"`
	Conn string `json:"connectionstring"`
}

func InitializeAPIHandlers() error {
	conf := new(DBlayerconfig)
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "./hydraweb/apiconfig.json")
	if err != nil {
		log.Println("Error decoding JSON", err)
		return err
	}
	h := newHydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Println("Error connection to db", err)
		return err
	}
	http.HandleFunc("/hydracrew/", h.handleHydraCrewRequests)
	return nil
}

func RunAPI() error {
	if err := InitializeAPIHandlers(); err != nil {
		return err
	}
	return http.ListenAndServe(":8061", nil)
}
