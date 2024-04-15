package auth_http

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	auth_dto "authorization/internal/domain/dto/auth"
)

type loginBodyReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type loginBodyResp struct {
	Token string `json:"token"`
}

// Login godoc
// @Tags         accounts
// @Description  Login to account
// @Accept       json
// @Produce      json
// @Param        User body loginBodyReq true "user data"
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

	var objBody loginBodyReq
	if err := json.Unmarshal(dataBody, &objBody); err != nil {
		http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
		return
	}

	// TODO Validate

	token, err := h.authHandler.Login(ctx, auth_dto.LoginDTO{
		Email:    objBody.Email,
		Password: objBody.Password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("Login error: %v", err), http.StatusBadRequest)
		return
	}

	resp := loginBodyResp{
		Token: token,
	}

	respBody, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Invalid response JSON body", http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
}
