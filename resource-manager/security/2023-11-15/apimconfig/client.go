package apimconfig

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type APIMConfigClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAPIMConfigClientWithBaseURI(endpoint string) APIMConfigClient {
	return APIMConfigClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
