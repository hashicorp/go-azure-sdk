package managedvmsizes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedVMSizesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewManagedVMSizesClientWithBaseURI(endpoint string) ManagedVMSizesClient {
	return ManagedVMSizesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
