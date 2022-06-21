package virtualmachinesizes

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineSizesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVirtualMachineSizesClientWithBaseURI(endpoint string) VirtualMachineSizesClient {
	return VirtualMachineSizesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
