package virtualmachinescalesetextensions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetExtensionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVirtualMachineScaleSetExtensionsClientWithBaseURI(endpoint string) VirtualMachineScaleSetExtensionsClient {
	return VirtualMachineScaleSetExtensionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
