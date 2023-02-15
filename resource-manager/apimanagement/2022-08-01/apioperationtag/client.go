package apioperationtag

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiOperationTagClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiOperationTagClientWithBaseURI(endpoint string) ApiOperationTagClient {
	return ApiOperationTagClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
