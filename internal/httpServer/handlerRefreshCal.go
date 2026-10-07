package httpserver

import (
	"encoding/json"
	"net/http"

	"wfm-json-to-ics/internal/getEventsJson"
)

func (cfg *LocalCfg) handlerRefreshCal(w http.ResponseWriter, req *http.Request) {
	type params struct {
		Args []string `json:"args"`
	}
	var p params
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode json", err)
		return
	}

	msg, err := getEventsJson.GetEventsJSON(p.Args, cfg.Config)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error(), err)
		return
	}
	type resp struct {
		Resp string `json:"resp"`
	}
	r := resp{
		Resp: msg,
	}
	respondWithJSON(w, http.StatusOK, r)
}
