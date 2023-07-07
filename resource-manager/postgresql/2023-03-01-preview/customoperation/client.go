package customoperation

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomOperationClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCustomOperationClientWithBaseURI(endpoint string) CustomOperationClient {
	return CustomOperationClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
