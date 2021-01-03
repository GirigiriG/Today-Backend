package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/GirigiriG/Clean-Architecture-golang/pkg/tools"

	"github.com/dgrijalva/jwt-go"
)

//IsAuthorized middleware
func IsAuthorized(next func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")

		tokens := r.Header["Token"]

		if len(tokens) == 0 {

			http.Redirect(w, r, "/", http.StatusForbidden)
			return
		}

		jwtToken, err := jwt.Parse(tokens[0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.Write(createErrorStruct(http.StatusForbidden, "Forbidden"))
				return nil, nil
			}
			return key, nil
		})

		if err == nil {
			if !jwtToken.Valid {
				w.Write(createErrorStruct(http.StatusForbidden, "Forbidden"))
				return
			}
			next(w, r)
		} else {
			w.Write(createErrorStruct(http.StatusForbidden, "Forbidden"))
		}
	})
}

var key []byte

//GenerateJWToken create jwt token
func GenerateJWToken() string {
	key = []byte(os.Getenv("P_KEY"))

	signer := jwt.New(jwt.SigningMethodHS512)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	signer.Claims = claims

	token, err := signer.SignedString(key)
	panicOnError(err)
	return token
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func createErrorStruct(code int, desc string) []byte {
	errStruct := tools.ErrorStatus{
		Code:        code,
		Description: desc,
	}

	resp, _ := json.Marshal(errStruct)
	return resp
}
