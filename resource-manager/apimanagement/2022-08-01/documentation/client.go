package documentation

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewDocumentationClientWithBaseURI(endpoint string) DocumentationClient {
	return DocumentationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
