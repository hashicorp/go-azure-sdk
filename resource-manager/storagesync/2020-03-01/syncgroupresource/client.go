package syncgroupresource

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupResourceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSyncGroupResourceClientWithBaseURI(endpoint string) SyncGroupResourceClient {
	return SyncGroupResourceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
