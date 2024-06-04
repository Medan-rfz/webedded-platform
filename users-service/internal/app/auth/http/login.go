package auth_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	auth_dto "webedded.users_management/internal/domain/dto/auth"
)

type loginBodyResp struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refreshToken"`
}

// Login godoc
// @Tags         accounts
// @Description  Login to account
// @Accept       json
// @Produce      json
// @Param        User body auth_dto.LoginDTO true "user data"
// @Success      200 {object} loginBodyResp
// @Failure      400
// @Failure      500
// @Router       /login [post]
func (h *authHttpHandler) Login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var err error
	var dataBody []byte
	if dataBody, err = io.ReadAll(r.Body); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	var data auth_dto.LoginDTO
	if err := json.Unmarshal(dataBody, &data); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	// TODO Validate

	loginResult, err := h.authHandler.Login(ctx, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Login error: %v", err), http.StatusBadRequest)
		return
	}

	resp := loginBodyResp{
		Token:        loginResult.Token,
		RefreshToken: loginResult.RefreshToken,
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Invalid response JSON body", http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
}
