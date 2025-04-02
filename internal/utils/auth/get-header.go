package authutil

import (
	"fmt"
	"net/http"
	"strings"
)

type Header struct {
	Request *http.Request
	Name    string
}

func GetTokenHeader(r *http.Request, header string) (string, error) {
	h := Header{
		Request: r,
		Name:    header,
	}

	token := h.getTokenFromHeader()

	if token == "" {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}

func (h *Header) getTokenFromHeader() string {
	session := h.Request.Header.Get(h.Name)
	arr := strings.Split(session, " ")

	if len(arr) != 2 || strings.ToUpper(arr[0]) != "BEARER" {
		return ""
	}

	return arr[1]
}
