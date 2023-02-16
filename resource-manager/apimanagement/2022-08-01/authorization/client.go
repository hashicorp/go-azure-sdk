package authorization

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAuthorizationClientWithBaseURI(endpoint string) AuthorizationClient {
	return AuthorizationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
