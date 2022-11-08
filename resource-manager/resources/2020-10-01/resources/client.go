package resources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewResourcesClientWithBaseURI(endpoint string) ResourcesClient {
	return ResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
