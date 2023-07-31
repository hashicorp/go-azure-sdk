package simgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SIMGroupsClient struct {
	Client *resourcemanager.Client
}

func NewSIMGroupsClientWithBaseURI(api sdkEnv.Api) (*SIMGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "simgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SIMGroupsClient: %+v", err)
	}

	return &SIMGroupsClient{
		Client: client,
	}, nil
}
