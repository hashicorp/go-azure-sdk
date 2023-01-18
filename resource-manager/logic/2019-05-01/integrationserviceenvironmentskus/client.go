package integrationserviceenvironmentskus

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSkusClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIntegrationServiceEnvironmentSkusClientWithBaseURI(endpoint string) IntegrationServiceEnvironmentSkusClient {
	return IntegrationServiceEnvironmentSkusClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
