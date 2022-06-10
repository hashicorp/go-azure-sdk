package operationsstatus

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationsStatusClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOperationsStatusClientWithBaseURI(endpoint string) OperationsStatusClient {
	return OperationsStatusClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
