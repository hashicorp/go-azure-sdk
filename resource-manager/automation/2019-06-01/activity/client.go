package activity

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityClient struct {
	Client  autorest.Client
	baseUri string
}

func NewActivityClientWithBaseURI(endpoint string) ActivityClient {
	return ActivityClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
