package resourcemanagementprivatelink

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceManagementPrivateLinkClient struct {
	Client  autorest.Client
	baseUri string
}

func NewResourceManagementPrivateLinkClientWithBaseURI(endpoint string) ResourceManagementPrivateLinkClient {
	return ResourceManagementPrivateLinkClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
