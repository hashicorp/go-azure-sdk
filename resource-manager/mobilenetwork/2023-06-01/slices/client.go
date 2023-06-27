package slices

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SlicesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSlicesClientWithBaseURI(endpoint string) SlicesClient {
	return SlicesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
