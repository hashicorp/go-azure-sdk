package usages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewUsagesClientWithBaseURI(endpoint string) UsagesClient {
	return UsagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
