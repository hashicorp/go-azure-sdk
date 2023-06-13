package crrjobdetails

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrrJobDetailsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCrrJobDetailsClientWithBaseURI(endpoint string) CrrJobDetailsClient {
	return CrrJobDetailsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
