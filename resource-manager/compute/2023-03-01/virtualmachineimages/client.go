package virtualmachineimages

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineImagesClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVirtualMachineImagesClientWithBaseURI(endpoint string) VirtualMachineImagesClient {
	return VirtualMachineImagesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
