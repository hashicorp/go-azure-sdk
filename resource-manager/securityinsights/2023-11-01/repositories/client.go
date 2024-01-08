package repositories

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RepositoriesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRepositoriesClientWithBaseURI(endpoint string) RepositoriesClient {
	return RepositoriesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
