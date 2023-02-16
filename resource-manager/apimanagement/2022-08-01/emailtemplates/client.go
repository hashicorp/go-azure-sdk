package emailtemplates

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplatesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEmailTemplatesClientWithBaseURI(endpoint string) EmailTemplatesClient {
	return EmailTemplatesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
