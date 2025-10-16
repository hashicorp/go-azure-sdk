package protectionintent

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectionIntentClient struct {
	Client  autorest.Client
	baseUri string
}

func NewProtectionIntentClientWithBaseURI(endpoint string) ProtectionIntentClient {
	return ProtectionIntentClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
