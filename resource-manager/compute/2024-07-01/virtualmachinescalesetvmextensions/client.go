package virtualmachinescalesetvmextensions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetVMExtensionsClient struct {
	Client *resourcemanager.Client
}

func NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualMachineScaleSetVMExtensionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "virtualmachinescalesetvmextensions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualMachineScaleSetVMExtensionsClient: %+v", err)
	}

	return &VirtualMachineScaleSetVMExtensionsClient{
		Client: client,
	}, nil
}
