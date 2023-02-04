package managementlocks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementLocksClient struct {
	Client *resourcemanager.Client
}

func NewManagementLocksClientWithBaseURI(api environments.Api) (*ManagementLocksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managementlocks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagementLocksClient: %+v", err)
	}

	return &ManagementLocksClient{
		Client: client,
	}, nil
}
