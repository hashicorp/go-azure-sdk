package powerbidedicated

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerBIDedicatedClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPowerBIDedicatedClientWithBaseURI(endpoint string) PowerBIDedicatedClient {
	return PowerBIDedicatedClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
