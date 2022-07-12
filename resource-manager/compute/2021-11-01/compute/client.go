package compute

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeClient struct {
	Client  autorest.Client
	baseUri string
}

func NewComputeClientWithBaseURI(endpoint string) ComputeClient {
	return ComputeClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
