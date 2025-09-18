package ver1

import "net/http"

func (h *Handlers) Greeting(w http.ResponseWriter, r *http.Request) {
	successResponse(w, http.StatusOK, "Hi there!")
}
