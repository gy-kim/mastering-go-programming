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

	fmt.Println("=======================")
	sbyte := []byte(`{"ShipID":1,"ShipClass":"Fighter","Captain":{"name":"Jaro","clearance level":10,"access codes":["ADA","LDL"]}}`)
	sinfo := new(ShipInfo)

	json.Unmarshal(sbyte, sinfo)
	fmt.Println(sinfo.ShipID, sinfo.ShipClass, sinfo.Captain.Name)

	ma := make(map[int]string)
	data := []byte(`{"1":"item1","2":"item2"}`)
	json.Unmarshal(data, &ma)
	fmt.Println(ma)
}
