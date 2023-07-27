package groupuser

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupUserClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGroupUserClientWithBaseURI(endpoint string) GroupUserClient {
	return GroupUserClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
