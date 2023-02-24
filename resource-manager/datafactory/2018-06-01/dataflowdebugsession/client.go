package dataflowdebugsession

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugSessionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDataFlowDebugSessionClientWithBaseURI(endpoint string) DataFlowDebugSessionClient {
	return DataFlowDebugSessionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
