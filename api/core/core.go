package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type API struct {
}

func (api *API) RespondJSON(w http.ResponseWriter, data interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("api/core: error marshaling data: %w", err)
	}

	if _, err := w.Write(jsonData); err != nil {
		return fmt.Errorf("api/core: error writing data: %w", err)
	}

	return nil
}

type ErrorResponse struct {
	Msg string `json:"message"`
}

var (
	ErrInternalError = errors.New("Internal error")
)

func (api *API) RespondError(w http.ResponseWriter, err error, status int) error {
	switch status {
	case http.StatusInternalServerError:
		fmt.Printf("api error: %s\n", err.Error())
		err = ErrInternalError
	}

	w.WriteHeader(status)
	return api.RespondJSON(w, ErrorResponse{Msg: err.Error()})
}
