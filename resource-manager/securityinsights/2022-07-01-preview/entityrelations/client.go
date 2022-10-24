package entityrelations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityRelationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEntityRelationsClientWithBaseURI(endpoint string) EntityRelationsClient {
	return EntityRelationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
