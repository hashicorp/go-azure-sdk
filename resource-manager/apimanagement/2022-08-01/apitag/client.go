package apitag

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiTagClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiTagClientWithBaseURI(endpoint string) ApiTagClient {
	return ApiTagClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
