package write

import (
	"encoding/json"
	"fmt"
	"github.com/cemayan/searchengine/common"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/service"
	"github.com/cemayan/searchengine/types"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (srv *Server) PostSelection(w http.ResponseWriter, r *http.Request) {
	rec := types.SelectionRequest{}
	err := json.NewDecoder(r.Body).Decode(&rec)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	svc := service.NewWriteService(constants.WriteApi)
	err = svc.Selection(rec)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	msg := fmt.Sprintf("record %s selected", rec.SelectedKey)
	apiResp := types.ApiResponse{Msg: &msg}

	if ok := json.NewEncoder(w).Encode(apiResp); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("json encode failed"))
		return
	}

}
