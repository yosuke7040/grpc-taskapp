package domain

type ErrValidationFailed struct {
	Msg string
}

func (e *ErrValidationFailed) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "failed to validate parameter"
}

type ErrNotFound struct {
	Msg string
}

func (e *ErrNotFound) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "not found"
}

type ErrQueryFailed struct {
	Msg string
}

func (e *ErrQueryFailed) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "failed to query"
}

type ErrPermissionDenied struct {
	Msg string
}

func (e *ErrPermissionDenied) Error() string {
	if e.Msg != "" {
		return e.Msg
	}
	return "permission denied"
}
