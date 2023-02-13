package suppressions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SuppressionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewSuppressionsClientWithBaseURI(endpoint string) SuppressionsClient {
	return SuppressionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
