package standbyvirtualmachines

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandbyVirtualMachinesClient struct {
	Client *resourcemanager.Client
}

func NewStandbyVirtualMachinesClientWithBaseURI(sdkApi sdkEnv.Api) (*StandbyVirtualMachinesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "standbyvirtualmachines", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StandbyVirtualMachinesClient: %+v", err)
	}

	return &StandbyVirtualMachinesClient{
		Client: client,
	}, nil
}
