package applyupdate

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyUpdateClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApplyUpdateClientWithBaseURI(endpoint string) ApplyUpdateClient {
	return ApplyUpdateClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
