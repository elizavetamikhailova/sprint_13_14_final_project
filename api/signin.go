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

	if password.Password != savedPassword {
		writeJSON(w, SignInResponse{Error: "Неверный пароль"})
		return
	}
	
	writeJSON(w, SignInResponse{Token: getFakeJwt(password.Password)})
}

func auth(next http.HandlerFunc) http.HandlerFunc {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        pass := os.Getenv("TODO_PASSWORD")
	
        if len(pass) > 0 {
            var jwt string 
            cookie, err := r.Cookie("token")
            if err == nil {
                jwt = cookie.Value
            }
        
			valid := isFakeJwtValid(jwt)


            if !valid {
                http.Error(w, "Authentification required", http.StatusUnauthorized)
                return
            }
        }
        next(w, r)
    })
}

func getFakeJwt(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	hashedBytes := hash.Sum(nil)
	return hex.EncodeToString(hashedBytes)
}

func isFakeJwtValid(jwt string) bool {
	return jwt != ""
}