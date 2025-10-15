package protectionintentresources

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntentResourcesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProtectionIntentResourcesClientWithBaseURI(endpoint string) ProtectionIntentResourcesClient {
	return ProtectionIntentResourcesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
