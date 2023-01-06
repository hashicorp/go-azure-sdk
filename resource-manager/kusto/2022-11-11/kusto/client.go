package kusto

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KustoClient struct {
	Client  autorest.Client
	baseUri string
}

func NewKustoClientWithBaseURI(endpoint string) KustoClient {
	return KustoClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
