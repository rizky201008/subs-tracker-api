package web

type responseCode int

const (
	RcSuccess      responseCode = 0
	RcError        responseCode = 1
	RcNotFound     responseCode = 2
	RcInvalid      responseCode = 3
	RcUnauthorized responseCode = 4
)
