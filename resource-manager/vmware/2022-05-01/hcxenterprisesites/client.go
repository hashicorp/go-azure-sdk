package hcxenterprisesites

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HcxEnterpriseSitesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewHcxEnterpriseSitesClientWithBaseURI(endpoint string) HcxEnterpriseSitesClient {
	return HcxEnterpriseSitesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
