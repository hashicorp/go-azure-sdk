package useridentity

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserIdentityClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUserIdentityClientWithBaseURI(endpoint string) UserIdentityClient {
	return UserIdentityClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
