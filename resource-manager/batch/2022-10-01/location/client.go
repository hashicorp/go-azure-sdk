package location

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewLocationClientWithBaseURI(endpoint string) LocationClient {
	return LocationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
