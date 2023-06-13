package aadproperties

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadPropertiesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewAadPropertiesClientWithBaseURI(endpoint string) AadPropertiesClient {
	return AadPropertiesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
