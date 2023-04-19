package codecontainer

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CodeContainerClient struct {
	Client  autorest.Client
	baseUri string
}

func NewCodeContainerClientWithBaseURI(endpoint string) CodeContainerClient {
	return CodeContainerClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
