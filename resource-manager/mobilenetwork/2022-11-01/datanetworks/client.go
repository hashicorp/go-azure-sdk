package datanetworks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataNetworksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDataNetworksClientWithBaseURI(endpoint string) DataNetworksClient {
	return DataNetworksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
