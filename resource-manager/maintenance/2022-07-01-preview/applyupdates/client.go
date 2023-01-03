package applyupdates

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplyUpdatesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewApplyUpdatesClientWithBaseURI(endpoint string) ApplyUpdatesClient {
	return ApplyUpdatesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
