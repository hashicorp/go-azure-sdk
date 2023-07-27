package networkstatus

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkStatusClient struct {
	Client  autorest.Client
	baseUri string
}

func NewNetworkStatusClientWithBaseURI(endpoint string) NetworkStatusClient {
	return NetworkStatusClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
