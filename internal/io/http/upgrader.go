package http

import (
	"net/http"

	"braces.dev/errtrace"
	ws "github.com/gorilla/websocket"
	"github.com/jarymor-ux/fer/internal/io/http/request"
	"github.com/jarymor-ux/fer/internal/io/http/response"
	"github.com/labstack/echo/v5"
)

func (h *Handlers) UpToWs(c *echo.Context) error {
	//TODO:тут будет вход в хаб и рега пользака
	return nil
}

func (h *Handlers) setUpAndUpgade( //nolint:unused
	w http.ResponseWriter,
	r *http.Request,
) (*ws.Conn, *request.Request, error) {
	req := request.NewRequest(w, r)

	conn, uerr := h.upgrader.Upgrade(w, r, nil)
	if uerr != nil {
		resp, err := response.NewHTTPError("Failed to create connection", http.StatusInternalServerError).
			GetJSONBytes()
		if err != nil {
			return nil, nil, errtrace.Wrap(err)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)

		if _, err = w.Write(resp); err != nil {
			return nil, nil, errtrace.Wrap(err)
		}

		return nil, nil, errtrace.Wrap(uerr)
	}
	// TODO:Выше нужно собирать данные пользака по типу айпи и тд
	return conn, &req, nil
}
