package managedclusterversion

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterVersionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewManagedClusterVersionClientWithBaseURI(endpoint string) ManagedClusterVersionClient {
	return ManagedClusterVersionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
