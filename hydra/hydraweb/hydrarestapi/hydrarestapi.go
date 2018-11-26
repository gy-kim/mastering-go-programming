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

func InitializeAPIHandlers() {
	conf := new(DBlayerconfig)
	err := hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, conf, "./hydraweb/apiconfig.json")
	if err != nil {
		log.Fatal("Error decoding JSON", err)
	}
	h := newHydraCrewReqHandler()
	err = h.connect(conf.DB, conf.Conn)
	if err != nil {
		log.Fatal("Error connection to db", err)
	}
	http.HandleFunc("/hydracrew/", h.handleHydraCrewRequests)
}

func RunAPI() {
	InitializeAPIHandlers()
	http.ListenAndServe(":8061", nil)
}
