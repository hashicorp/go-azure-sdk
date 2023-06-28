package simgroups

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMGroupsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSIMGroupsClientWithBaseURI(endpoint string) SIMGroupsClient {
	return SIMGroupsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
