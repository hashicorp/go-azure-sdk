package managementgroups

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementGroupsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewManagementGroupsClientWithBaseURI(endpoint string) ManagementGroupsClient {
	return ManagementGroupsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
