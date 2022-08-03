package virtualmachines

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachinesClient struct {
	Client  autorest.Client
	Client2 *resourcemanager.Client
	baseUri string
}

func NewVirtualMachinesClientWithBaseURI(endpoint string) VirtualMachinesClient {
	return VirtualMachinesClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		Client2: resourcemanager.NewResourceManagerClient(environments.ApiEndpoint(endpoint)),
		baseUri: endpoint,
	}
}
