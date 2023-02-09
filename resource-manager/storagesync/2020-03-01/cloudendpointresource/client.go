package cloudendpointresource

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointResourceClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCloudEndpointResourceClientWithBaseURI(endpoint string) CloudEndpointResourceClient {
	return CloudEndpointResourceClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
