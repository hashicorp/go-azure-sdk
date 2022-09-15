package virtualmachineruncommands

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineRunCommandsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewVirtualMachineRunCommandsClientWithBaseURI(endpoint string) VirtualMachineRunCommandsClient {
	return VirtualMachineRunCommandsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
