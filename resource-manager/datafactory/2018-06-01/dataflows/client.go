package dataflows

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDataFlowsClientWithBaseURI(endpoint string) DataFlowsClient {
	return DataFlowsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
