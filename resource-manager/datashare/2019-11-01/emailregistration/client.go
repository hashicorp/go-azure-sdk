package emailregistration

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailRegistrationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEmailRegistrationClientWithBaseURI(endpoint string) EmailRegistrationClient {
	return EmailRegistrationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
