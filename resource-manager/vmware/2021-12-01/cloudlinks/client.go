package cloudlinks

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudLinksClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCloudLinksClientWithBaseURI(endpoint string) CloudLinksClient {
	return CloudLinksClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
