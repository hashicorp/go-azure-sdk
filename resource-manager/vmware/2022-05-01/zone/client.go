package zone

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZoneClient struct {
	Client  autorest.Client
	baseUri string
}

func NewZoneClientWithBaseURI(endpoint string) ZoneClient {
	return ZoneClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
