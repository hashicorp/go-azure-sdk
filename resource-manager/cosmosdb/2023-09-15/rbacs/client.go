package rbacs

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RbacsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRbacsClientWithBaseURI(endpoint string) RbacsClient {
	return RbacsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
