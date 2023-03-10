package securitypins

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPINsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSecurityPINsClientWithBaseURI(endpoint string) SecurityPINsClient {
	return SecurityPINsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
