package dedicatedhost

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DedicatedHostClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDedicatedHostClientWithBaseURI(endpoint string) DedicatedHostClient {
	return DedicatedHostClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
