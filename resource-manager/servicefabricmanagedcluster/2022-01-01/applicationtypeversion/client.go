package applicationtypeversion

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationTypeVersionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApplicationTypeVersionClientWithBaseURI(endpoint string) ApplicationTypeVersionClient {
	return ApplicationTypeVersionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
