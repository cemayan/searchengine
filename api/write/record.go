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

func (srv *Server) PostRecord(w http.ResponseWriter, r *http.Request) {

	rec := types.RecordRequest{}
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		logrus.Error(err)
		return
	}

	svc := service.NewWriteService(constants.WriteApi)
	svc.Start(rec.Data)

	msg := fmt.Sprintf("%s record added to database successfully", rec.Data)
	apiResp := types.ApiResponse{Msg: &msg}

	w.Header().Set("Content-Type", "application/json")

	if ok := json.NewEncoder(w).Encode(apiResp); ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(common.GetError("json encode failed"))
	}

}
