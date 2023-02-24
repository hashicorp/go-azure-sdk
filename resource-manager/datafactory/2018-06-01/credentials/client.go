package credentials

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCredentialsClientWithBaseURI(endpoint string) CredentialsClient {
	return CredentialsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
