package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New(validator.WithRequiredStructEnabled())
}

func BindAndValidate(data interface{}, r *http.Request) error {
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	defer r.Body.Close()

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(data); err != nil {
		return fmt.Errorf("invalid request body: %w", err)
	}

	if dec.More() {
		return fmt.Errorf("unexpected additional JSON in body")
	}

	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}

type FindServiceRequest struct {
	Intent string `json:"intent" validate:"required"`
}
type ServiceData struct {
	ServiceID   int    `json:"service_id"`
	ServiceName string `json:"service_name"`
}
type Message struct {
	Success bool        `json:"success"`
	Data    ServiceData `json:"data"`
	Error   string      `json:"error"`
}
