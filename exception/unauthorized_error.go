package exception

type UnauthorizedError struct {
	Err string
}

func (e UnauthorizedError) Error() string {
	return e.Err
}
