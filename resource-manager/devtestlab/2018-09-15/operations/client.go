package operations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOperationsClientWithBaseURI(endpoint string) OperationsClient {
	return OperationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
