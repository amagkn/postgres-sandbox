package ver1

import (
	"context"
	"net/http"

	"github.com/amagkn/postgres-sandbox/internal/air/dto"
	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/amagkn/postgres-sandbox/pkg/logger"
	"github.com/amagkn/postgres-sandbox/pkg/validation"
)

func (h *Handlers) TripAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	input := dto.TripAllInput{
		Town: r.URL.Query().Get("town"),
	}

	invalidFields, err := validation.ValidateStruct(&input)
	if err != nil {
		logger.Error(err, "validation.ValidateStruct")
		errorResponse(w, http.StatusBadRequest, errorPayload{Type: base_errors.Validation, Details: invalidFields})
		return
	}

	output, err := h.uc.TripAll(ctx, input)
	if err != nil {
		logger.Error(err, "uc.TripAll")
		errorResponse(w, http.StatusInternalServerError, errorPayload{Type: base_errors.InternalServer})
		return
	}

	successResponse(w, http.StatusOK, output)
}
