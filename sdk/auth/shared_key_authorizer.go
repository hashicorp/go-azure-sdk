package auth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

var _ Authorizer = &SharedKeyAuthorizer{}

type SharedKeyAuthorizer struct {
}

func (s SharedKeyAuthorizer) Token(ctx context.Context, request *http.Request) (*oauth2.Token, error) {
	return nil, fmt.Errorf("internal-error: SharedKeyAuthorizer is not implemented")
}

func (s SharedKeyAuthorizer) AuxiliaryTokens(ctx context.Context, request *http.Request) ([]*oauth2.Token, error) {
	return nil, fmt.Errorf("internal-error: SharedKeyAuthorizer is not implemented")
}
