package usersession

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSessionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUserSessionClientWithBaseURI(endpoint string) UserSessionClient {
	return UserSessionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
