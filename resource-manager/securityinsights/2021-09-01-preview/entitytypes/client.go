package entitytypes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityTypesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEntityTypesClientWithBaseURI(endpoint string) EntityTypesClient {
	return EntityTypesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
