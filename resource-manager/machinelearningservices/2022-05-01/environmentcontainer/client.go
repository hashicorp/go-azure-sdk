package environmentcontainer

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentContainerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEnvironmentContainerClientWithBaseURI(endpoint string) EnvironmentContainerClient {
	return EnvironmentContainerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
