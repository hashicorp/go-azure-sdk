package modelcontainer

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModelContainerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewModelContainerClientWithBaseURI(endpoint string) ModelContainerClient {
	return ModelContainerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
