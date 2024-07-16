package read

import (
	"encoding/json"
	"fmt"
	"github.com/cemayan/searchengine/common"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/service"
	"net/http"
)

func (srv *Server) GetResults(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	query := r.Header.Get(constants.XSearchEngineQuery)

	if query != "" {

		svc := service.NewReadService(constants.ReadApi)
		resp := svc.GetResults(query)

		if ok := json.NewEncoder(w).Encode(resp); ok != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(common.GetError("json encode failed"))
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError(fmt.Sprintf("%s cannot be empty", constants.XSearchEngineQuery)))
		return
	}
}
