package offboardfromd4api

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OffboardFromD4APIClient struct {
	Client  autorest.Client
	baseUri string
}

func NewOffboardFromD4APIClientWithBaseURI(endpoint string) OffboardFromD4APIClient {
	return OffboardFromD4APIClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
