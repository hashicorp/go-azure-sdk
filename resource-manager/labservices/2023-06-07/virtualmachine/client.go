package virtualmachine

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineClient struct {
	Client *resourcemanager.Client
}

func NewVirtualMachineClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualMachineClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "virtualmachine", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualMachineClient: %+v", err)
	}

	return &VirtualMachineClient{
		Client: client,
	}, nil
}
