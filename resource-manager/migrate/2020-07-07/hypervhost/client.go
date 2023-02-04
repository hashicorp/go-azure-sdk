package hypervhost

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVHostClient struct {
	Client  autorest.Client
	baseUri string
}

func NewHyperVHostClientWithBaseURI(endpoint string) HyperVHostClient {
	return HyperVHostClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
