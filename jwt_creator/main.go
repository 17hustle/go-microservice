package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	jwt "github.com/golang-jwt/jwt/v5"
)

var SignInKey = []byte(os.Getenv("SECRET_KEY"))

func GetJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = "hustler"
	claims["aud"] = "bill.jwt.so"
	claims["iss"] = "jwt.so"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(SignInKey)
	if err != nil {
		return "", fmt.Errorf("Something went wrong: %v", err)
	}
	return tokenString, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	validToken, err := GetJWT()
	if err != nil {
		fmt.Println("Failed to generate the token:", err)
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, validToken)
}

func handleRequests() {
	http.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("Server is starting on port 8080")
	handleRequests()
}
