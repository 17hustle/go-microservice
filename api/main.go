package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

var myKey = []byte(os.Getenv("SECRET_KEY")) // Retrieve the secret key from an environment variable

func handleRequests() {
	http.Handle("/", isAuthorized(homepage)) // Set up route handler
	log.Fatal(http.ListenAndServe(":9002", nil)) // Start server on port 9002
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Secret information")
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if tokenString := r.Header.Get("Token"); tokenString != "" {
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				// Ensure the token's method is HMAC
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Invalid method")
				}
				return myKey, nil
			})

			if err != nil || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Assertions for claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
				return
			}

			aud := "bill.jwt.so"
			if claimAud, ok := claims["aud"].(string); !ok || claimAud != aud {
				http.Error(w, "Invalid audience", http.StatusUnauthorized)
				return
			}

			iss := "jwt.so"
			if claimIss, ok := claims["iss"].(string); !ok || claimIss != iss {
				http.Error(w, "Invalid issuer", http.StatusUnauthorized)
				return
			}

			// Call the actual endpoint handler if authorized
			endpoint(w, r)
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}


func main() {
	fmt.Println("Server is starting")
	handleRequests()
}
