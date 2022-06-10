package applicationtype

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTypeClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApplicationTypeClientWithBaseURI(endpoint string) ApplicationTypeClient {
	return ApplicationTypeClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
