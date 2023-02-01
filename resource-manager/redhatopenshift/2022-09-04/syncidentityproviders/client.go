package syncidentityproviders

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncIdentityProvidersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSyncIdentityProvidersClientWithBaseURI(endpoint string) SyncIdentityProvidersClient {
	return SyncIdentityProvidersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
