package errapp

import "errors"

var (
	ErrNoDocuments = errors.New("Documento não encontrado")
)
