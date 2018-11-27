package hydrarestapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gy-kim/mastering-go-programming/hydra/hydradblayer"
)

type hydracrewreqhandler struct {
	dbConn hydradblayer.DBLayer
}

func newHydraCrewReqHandler() *hydracrewreqhandler {
	return new(hydracrewreqhandler)
}

func (hcwerq *hydracrewreqhandler) connect(o, conn string) error {
	dblayer, err := hydradblayer.ConnectDatabase(o, conn)
	if err != nil {
		return err
	}
	hcwerq.dbConn = dblayer
	return nil
}

func (hcwreq *hydracrewreqhandler) handleHydraCrewRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ids := r.RequestURI[len("/hydracrew/"):]
		// /hydracrew/2
		id, err := strconv.Atoi(ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "id %s provided is not of valid member.\n", ids)
			return
		}
		cm, err := hcwreq.dbConn.FindMember(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occured while searching for id %d\n", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&cm)
	case "POST":
		cm := new(hydradblayer.CrewMember)
		err := json.NewDecoder(r.Body).Decode(cm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error %s occured", err)
			return
		}
		err = hcwreq.dbConn.AddMember(cm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occured while adding a crew member to the Hydra database", err)
			return
		}
		fmt.Fprintf(w, "Successfully interted id %d\n", cm.ID)
	}
}
