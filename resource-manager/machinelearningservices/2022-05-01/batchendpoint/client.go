package batchendpoint

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchEndpointClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBatchEndpointClientWithBaseURI(endpoint string) BatchEndpointClient {
	return BatchEndpointClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
