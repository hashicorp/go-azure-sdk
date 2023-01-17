package costs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCostsClientWithBaseURI(endpoint string) CostsClient {
	return CostsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
