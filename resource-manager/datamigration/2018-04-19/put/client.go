package put

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PUTClient struct {
	Client  autorest.Client
	baseUri string
}

func NewPUTClientWithBaseURI(endpoint string) PUTClient {
	return PUTClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
