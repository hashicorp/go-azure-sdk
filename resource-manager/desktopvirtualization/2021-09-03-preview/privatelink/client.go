package privatelink

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPrivateLinkClientWithBaseURI(endpoint string) PrivateLinkClient {
	return PrivateLinkClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
