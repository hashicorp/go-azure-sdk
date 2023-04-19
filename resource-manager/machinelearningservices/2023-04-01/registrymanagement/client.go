package registrymanagement

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryManagementClient struct {
	Client  autorest.Client
	baseUri string
}

func NewRegistryManagementClientWithBaseURI(endpoint string) RegistryManagementClient {
	return RegistryManagementClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
