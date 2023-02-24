package factories

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FactoriesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewFactoriesClientWithBaseURI(endpoint string) FactoriesClient {
	return FactoriesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
