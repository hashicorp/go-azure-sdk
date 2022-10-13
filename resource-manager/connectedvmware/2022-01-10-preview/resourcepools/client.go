package resourcepools

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcePoolsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewResourcePoolsClientWithBaseURI(endpoint string) ResourcePoolsClient {
	return ResourcePoolsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
