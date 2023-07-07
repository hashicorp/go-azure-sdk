package migrations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMigrationsClientWithBaseURI(endpoint string) MigrationsClient {
	return MigrationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
