package hcireports

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HCIReportsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewHCIReportsClientWithBaseURI(endpoint string) HCIReportsClient {
	return HCIReportsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
