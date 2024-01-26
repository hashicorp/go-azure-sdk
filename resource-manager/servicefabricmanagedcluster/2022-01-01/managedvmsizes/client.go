package managedvmsizes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedVMSizesClient struct {
	Client *resourcemanager.Client
}

func NewManagedVMSizesClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedVMSizesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managedvmsizes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedVMSizesClient: %+v", err)
	}

	return &ManagedVMSizesClient{
		Client: client,
	}, nil
}
