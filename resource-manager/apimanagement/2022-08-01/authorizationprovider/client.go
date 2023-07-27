package authorizationprovider

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationProviderClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAuthorizationProviderClientWithBaseURI(endpoint string) AuthorizationProviderClient {
	return AuthorizationProviderClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
