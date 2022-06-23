package keyvalues

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValuesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewKeyValuesClientWithBaseURI(endpoint string) KeyValuesClient {
	return KeyValuesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
