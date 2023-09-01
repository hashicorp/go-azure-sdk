package mhsmprivatelinkresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMPrivateLinkResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewMHSMPrivateLinkResourcesClientWithBaseURI(endpoint string) MHSMPrivateLinkResourcesClient {
	return MHSMPrivateLinkResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
