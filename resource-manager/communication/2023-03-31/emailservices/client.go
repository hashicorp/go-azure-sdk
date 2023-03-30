package emailservices

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailServicesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEmailServicesClientWithBaseURI(endpoint string) EmailServicesClient {
	return EmailServicesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
