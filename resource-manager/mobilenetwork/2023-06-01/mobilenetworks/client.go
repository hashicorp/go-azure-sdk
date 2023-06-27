package mobilenetworks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileNetworksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMobileNetworksClientWithBaseURI(endpoint string) MobileNetworksClient {
	return MobileNetworksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
