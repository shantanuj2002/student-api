package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/shantanuj2002/students-api/internal/types"
	"github.com/shantanuj2002/students-api/internal/utils/response"
)

var validate = validator.New()

func New() http.HandlerFunc {
	return func(w http.ResponseWriter,
		r *http.Request) {
		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty body")))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		// Validate the decoded student
		if err := validate.Struct(student); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			errorMessages := make(map[string]string)
			for _, ve := range validationErrors {
				errorMessages[ve.Field()] = fmt.Sprintf("failed on the '%s' tag", ve.Tag())
			}

			response.WriteJson(w, http.StatusBadRequest, map[string]interface{}{
				"status":  "Error",
				"message": "Validation failed",
				"errors":  errorMessages,
			})
			return
		}

		slog.Info("Creating a student")

		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
