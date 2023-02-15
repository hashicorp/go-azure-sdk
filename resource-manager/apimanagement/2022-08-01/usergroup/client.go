package usergroup

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserGroupClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUserGroupClientWithBaseURI(endpoint string) UserGroupClient {
	return UserGroupClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
