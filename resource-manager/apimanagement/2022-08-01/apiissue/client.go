package apiissue

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiIssueClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiIssueClientWithBaseURI(endpoint string) ApiIssueClient {
	return ApiIssueClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
