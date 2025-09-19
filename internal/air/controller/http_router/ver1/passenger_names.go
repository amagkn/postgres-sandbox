package ver1

import (
	"context"
	"net/http"

	"github.com/amagkn/postgres-sandbox/pkg/base_errors"
	"github.com/amagkn/postgres-sandbox/pkg/logger"
)

func (h *Handlers) PassengerNames(w http.ResponseWriter, _ *http.Request) {
	ctx := context.Background()

	output, err := h.uc.PassengerNames(ctx)
	if err != nil {
		logger.Error(err, "uc.PassengerNames")
		errorResponse(w, http.StatusInternalServerError, errorPayload{Type: base_errors.InternalServer})
		return
	}

	successResponse(w, http.StatusOK, output)
}
