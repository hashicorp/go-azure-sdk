package availableskus

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableSkusClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAvailableSkusClientWithBaseURI(endpoint string) AvailableSkusClient {
	return AvailableSkusClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
