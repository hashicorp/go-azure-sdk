package vmhhostlist

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMHHostListClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVMHHostListClientWithBaseURI(endpoint string) VMHHostListClient {
	return VMHHostListClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
