package apischema

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiSchemaClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiSchemaClientWithBaseURI(endpoint string) ApiSchemaClient {
	return ApiSchemaClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
