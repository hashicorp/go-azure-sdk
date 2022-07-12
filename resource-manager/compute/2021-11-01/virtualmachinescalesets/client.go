package virtualmachinescalesets

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVirtualMachineScaleSetsClientWithBaseURI(endpoint string) VirtualMachineScaleSetsClient {
	return VirtualMachineScaleSetsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
