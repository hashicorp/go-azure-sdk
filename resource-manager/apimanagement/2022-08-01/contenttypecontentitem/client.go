package contenttypecontentitem

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentTypeContentItemClient struct {
	Client  autorest.Client
	baseUri string
}

func NewContentTypeContentItemClientWithBaseURI(endpoint string) ContentTypeContentItemClient {
	return ContentTypeContentItemClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
