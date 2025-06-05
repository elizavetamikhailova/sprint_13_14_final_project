package api

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"os"
)

type Password struct {
	Password      string  `json:"password"`
}

type SignInResponse struct {
	Token    string  `json:"token,omitempty"`
	Error string `json:"error,omitempty"`
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	var password Password
	if err := json.NewDecoder(r.Body).Decode(&password); err != nil {
		writeJSON(w, SignInResponse{Error: "Invalid JSON format"})
		return
	}

	savedPassword := os.Getenv("TODO_PASSWORD")
	if savedPassword == "" {
		savedPassword = "12345"
	}

	if password.Password != savedPassword {
		writeJSON(w, SignInResponse{Error: "Неверный пароль"})
		return
	}

	hash := sha256.New()
	hash.Write([]byte(password.Password))
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)
	
	writeJSON(w, SignInResponse{Token: hashedString})
}