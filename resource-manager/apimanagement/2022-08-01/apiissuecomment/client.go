package apiissuecomment

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiIssueCommentClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApiIssueCommentClientWithBaseURI(endpoint string) ApiIssueCommentClient {
	return ApiIssueCommentClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
