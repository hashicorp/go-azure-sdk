package integrationruntimeobjectmetadata

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeObjectMetadataClient struct {
	Client  autorest.Client
	baseUri string
}

func NewIntegrationRuntimeObjectMetadataClientWithBaseURI(endpoint string) IntegrationRuntimeObjectMetadataClient {
	return IntegrationRuntimeObjectMetadataClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
