package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type CrewMember struct {
		ID                 int      `json:"id,omitempy"`
		Name               string   `json:"name,omitempty"`
		SecurityClearnance int      `json:"clearancelevel"`
		AccessCodes        []string `json:"accesscodes"`
	}

	type ShipInfo struct {
		ShipID    int
		ShipClass string
		Captain   CrewMember
	}

	cm := CrewMember{1, "Jaro", 10, []string{"ADA", "LAL"}}
	si := ShipInfo{1, "Fighter", cm}
	b, err := json.Marshal(si)

	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))

	m := map[string]int{"item1": 1, "item2": 2}
	b, err = json.Marshal(m)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(b))

	s := []CrewMember{cm, CrewMember{2, "Jim", 5, []string{"TLT", "RAR"}}}
	bSlice, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(string(bSlice))
}
