package auth_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	auth_dto "webedded.users_management/internal/domain/dto/auth"
)

// Refresh godoc
// @Tags         accounts
// @Description  Refresh jwt token
// @Accept       json
// @Produce      json
// @Param        User body auth_dto.RefreshDTO true "user data"
// @Success      200 {object} loginBodyResp
// @Failure      400
// @Failure      500
// @Router       /refresh [post]
func (h *authHttpHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var err error
	var dataBody []byte
	if dataBody, err = io.ReadAll(r.Body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	var objBody auth_dto.RefreshDTO
	if err := json.Unmarshal(dataBody, &objBody); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	// TODO Validate

	newToken, err := h.authHandler.Refresh(ctx, objBody)
	if err != nil {
		http.Error(w, fmt.Sprintf("Registration error: %v", err), http.StatusBadRequest)
	}

	resp := loginBodyResp{
		Token:        newToken,
		RefreshToken: objBody.RefreshToken,
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Invalid response JSON body", http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
}
