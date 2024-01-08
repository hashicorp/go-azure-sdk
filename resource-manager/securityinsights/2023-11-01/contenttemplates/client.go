package contenttemplates

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTemplatesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewContentTemplatesClientWithBaseURI(endpoint string) ContentTemplatesClient {
	return ContentTemplatesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
