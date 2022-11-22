package auth

import (
	"context"
	"golang.org/x/oauth2"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

// Authorizer is anything that can return an access token for authorizing API connections
type Authorizer interface {
	Token(ctx context.Context) (*oauth2.Token, error)

	AuxiliaryTokens(ctx context.Context) ([]*oauth2.Token, error)
}
