package hypervsites

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HyperVSitesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewHyperVSitesClientWithBaseURI(endpoint string) HyperVSitesClient {
	return HyperVSitesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
