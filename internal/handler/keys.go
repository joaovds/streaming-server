package handler

import "net/http"

type KeyHandler interface {
	Auth(w http.ResponseWriter, r *http.Request)
}

type keyHandler struct {
}

func NewKeyHandler() *keyHandler {
	return &keyHandler{}
}

func (k *keyHandler) Auth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
