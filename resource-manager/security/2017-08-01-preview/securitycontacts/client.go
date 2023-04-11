package securitycontacts

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContactsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSecurityContactsClientWithBaseURI(endpoint string) SecurityContactsClient {
	return SecurityContactsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
