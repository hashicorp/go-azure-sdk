package entities

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitiesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEntitiesClientWithBaseURI(endpoint string) EntitiesClient {
	return EntitiesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
