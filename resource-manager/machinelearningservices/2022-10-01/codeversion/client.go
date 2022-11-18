package codeversion

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CodeVersionClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCodeVersionClientWithBaseURI(endpoint string) CodeVersionClient {
	return CodeVersionClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
