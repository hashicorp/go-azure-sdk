package customlocations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomLocationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCustomLocationsClientWithBaseURI(endpoint string) CustomLocationsClient {
	return CustomLocationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
