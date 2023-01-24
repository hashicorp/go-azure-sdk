package fieresource

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FieResourceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewFieResourceClientWithBaseURI(endpoint string) FieResourceClient {
	return FieResourceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
