package auth

import "golang.org/x/oauth2"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

// Authorizer is anything that can return an access token for authorizing API connections
type Authorizer interface {
	Token() (*oauth2.Token, error)
	AuxiliaryTokens() ([]*oauth2.Token, error)
}
