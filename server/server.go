package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SouvenirHandler(w http.ResponseWriter, r *http.Request) {
	rawID := mux.Vars(r)["id"]
	id64, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	id := int(id64)
	id = id + 1
}
