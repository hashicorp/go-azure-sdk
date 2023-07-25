package global

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalClient struct {
	Client  autorest.Client
	baseUri string
}

func NewGlobalClientWithBaseURI(endpoint string) GlobalClient {
	return GlobalClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
