package usertoken

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTokenClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUserTokenClientWithBaseURI(endpoint string) UserTokenClient {
	return UserTokenClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
