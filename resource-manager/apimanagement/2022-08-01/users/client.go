package users

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUsersClientWithBaseURI(endpoint string) UsersClient {
	return UsersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
