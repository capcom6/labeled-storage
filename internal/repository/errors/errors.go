package errors

type errorConst string

const ErrKeyNotFound errorConst = "key not found"

func (e errorConst) Error() string {
	return string(e)
}
