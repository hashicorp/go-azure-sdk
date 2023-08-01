package vmhhostlist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMHHostListClient struct {
	Client *resourcemanager.Client
}

func NewVMHHostListClientWithBaseURI(sdkApi sdkEnv.Api) (*VMHHostListClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "vmhhostlist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VMHHostListClient: %+v", err)
	}

	return &VMHHostListClient{
		Client: client,
	}, nil
}
