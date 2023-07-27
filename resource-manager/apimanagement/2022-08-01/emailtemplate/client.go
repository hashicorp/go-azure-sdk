package emailtemplate

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailTemplateClient struct {
	Client  autorest.Client
	baseUri string
}

func NewEmailTemplateClientWithBaseURI(endpoint string) EmailTemplateClient {
	return EmailTemplateClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
