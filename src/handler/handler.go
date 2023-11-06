package handler

import "ggf/src/redirect"

type Handler struct {
	redirectStore redirect.Store
	validator     *Validator
}

func NewHandler(redirectStore redirect.Store) *Handler {
	v := NewValidator()
	return &Handler{
		redirectStore: redirectStore,
		validator:     v,
	}
}
