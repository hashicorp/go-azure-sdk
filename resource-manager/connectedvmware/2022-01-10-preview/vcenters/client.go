package vcenters

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCentersClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVCentersClientWithBaseURI(endpoint string) VCentersClient {
	return VCentersClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
