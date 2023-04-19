package componentcontainer

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentContainerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewComponentContainerClientWithBaseURI(endpoint string) ComponentContainerClient {
	return ComponentContainerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
