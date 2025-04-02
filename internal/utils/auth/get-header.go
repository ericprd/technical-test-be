package authutil

import (
	"fmt"
	"net/http"
	"strings"
)

func GetTokenHeader(r *http.Request) (string, error) {

	token := getTokenFromHeader(r)

	if token == "" {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}

func getTokenFromHeader(r *http.Request) string {
	session := r.Header.Get("Authorization")
	arr := strings.Split(session, " ")

	if len(arr) != 2 || strings.ToUpper(arr[0]) != "BEARER" {
		return ""
	}

	return arr[1]
}
