package save

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/dimsog/httpmocks-backend/internal/http-server/render"
	"github.com/go-playground/validator/v10"
)

type Request struct {
	ContentType string `json:"content-type" validate:"required"`
	Response    string `json:"response"`
}

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request Request
		validate := validator.New()

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Debug("Invalid request: " + err.Error())
		}

		err = validate.Struct(request)

		if err != nil {
			err := render.ValidationError(w, err)
			if err != nil {
				log.Error(err.Error())
			}
			return
		}

		render.Success(w)
	}
}
