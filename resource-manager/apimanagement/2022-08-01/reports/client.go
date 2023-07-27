package reports

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReportsClientWithBaseURI(endpoint string) ReportsClient {
	return ReportsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
