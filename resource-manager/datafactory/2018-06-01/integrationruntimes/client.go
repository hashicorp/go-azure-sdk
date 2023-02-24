package integrationruntimes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIntegrationRuntimesClientWithBaseURI(endpoint string) IntegrationRuntimesClient {
	return IntegrationRuntimesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
