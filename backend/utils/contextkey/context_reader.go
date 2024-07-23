package contextkey

import (
	"context"
	"errors"
)

type IContextReader interface {
	GetUserID(ctx context.Context) (string, error)
}

type ContextReader struct{}

func NewContextReader() IContextReader {
	return &ContextReader{}
}

func (c *ContextReader) GetUserID(ctx context.Context) (string, error) {
	v := ctx.Value(ContextKeyUserID)
	if v == nil {
		return "", errors.New("user id not found in context")
	}

	userID, ok := v.(string)
	if !ok {
		return "", errors.New("user id is not a string")
	}

	return userID, nil
}
