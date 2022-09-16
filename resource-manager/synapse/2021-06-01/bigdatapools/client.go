package bigdatapools

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BigDataPoolsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewBigDataPoolsClientWithBaseURI(endpoint string) BigDataPoolsClient {
	return BigDataPoolsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
