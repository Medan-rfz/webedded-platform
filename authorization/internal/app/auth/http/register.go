package auth_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	authdto "authorization/internal/domain/dto/auth"
)

// Register godoc
// @Tags         accounts
// @Description  Register of account
// @Accept       json
// @Produce      json
// @Param        User body authdto.RegisterDTO true "user data"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /register [post]
func (h *authHttpHandler) Register(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var err error
	var dataBody []byte
	if dataBody, err = io.ReadAll(r.Body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	var objBody authdto.RegisterDTO
	if err := json.Unmarshal(dataBody, &objBody); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	// TODO Validate

	if err := h.authHandler.Register(ctx, objBody); err != nil {
		http.Error(w, fmt.Sprintf("Registration error: %v", err), http.StatusBadRequest)
	}
}
