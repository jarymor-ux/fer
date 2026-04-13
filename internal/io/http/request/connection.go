package request

import (
	"encoding/json"
	"net"
	"net/http"

	"github.com/jarymor-ux/fer/internal/domain/types"
	"github.com/jarymor-ux/fer/internal/utils"
)

type Request struct {
	SenderID types.UserID
	Username string `json:"username"`
	IP       net.IP
}

func NewRequest(w http.ResponseWriter, r *http.Request) Request {
	req := new(Request)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	ip := net.IP(host) //WARN:Мб не сработает

	return Request{
		SenderID: types.UserID(utils.GenerateUUID()),
		Username: req.Username,
		IP:       ip,
	}
}
