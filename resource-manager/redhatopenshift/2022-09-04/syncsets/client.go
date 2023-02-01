package syncsets

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncSetsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSyncSetsClientWithBaseURI(endpoint string) SyncSetsClient {
	return SyncSetsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
