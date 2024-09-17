package managedclusterversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterVersionClient struct {
	Client *resourcemanager.Client
}

func NewManagedClusterVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedClusterVersionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedclusterversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedClusterVersionClient: %+v", err)
	}

	return &ManagedClusterVersionClient{
		Client: client,
	}, nil
}
