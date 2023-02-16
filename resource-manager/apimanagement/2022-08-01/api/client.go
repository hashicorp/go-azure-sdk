package api

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiClientWithBaseURI(endpoint string) ApiClient {
	return ApiClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
