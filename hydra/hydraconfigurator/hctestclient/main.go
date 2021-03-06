package main

import (
	"fmt"

	"github.com/gy-kim/mastering-go-programming/hydra/hydraconfigurator"
)

type ConfS struct {
	TS      string  `name:"testString" json:"testString" xml:"testString"`
	TB      bool    `name:"testBool" json:"testBool" xml:"testBool"`
	TF      float64 `name:"testFloat" json:"testFloat" xml:"testFloat"`
	TestInt int
}

func main() {
	configstruct := new(ConfS)

	// hydraconfigurator.GetConfiguration(hydraconfigurator.CUSTOM, configstruct, "configfile.conf")
	// hydraconfigurator.GetConfiguration(hydraconfigurator.JSON, configstruct, "configfile.json")
	hydraconfigurator.GetConfiguration(hydraconfigurator.XML, configstruct, "configfile.xml")
	fmt.Println(*configstruct)

	if configstruct.TB {
		fmt.Println("bool is true")
	}

	fmt.Println(float64(4.8 * configstruct.TF))
	fmt.Println(5 * configstruct.TestInt)
	fmt.Println(configstruct.TS)
}
