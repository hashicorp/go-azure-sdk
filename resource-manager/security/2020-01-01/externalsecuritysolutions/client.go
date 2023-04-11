package externalsecuritysolutions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalSecuritySolutionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewExternalSecuritySolutionsClientWithBaseURI(endpoint string) ExternalSecuritySolutionsClient {
	return ExternalSecuritySolutionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
