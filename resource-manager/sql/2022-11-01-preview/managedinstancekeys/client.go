package managedinstancekeys

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceKeysClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstanceKeysClientWithBaseURI(api environments.Api) (*ManagedInstanceKeysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "managedinstancekeys", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstanceKeysClient: %+v", err)
	}

	return &ManagedInstanceKeysClient{
		Client: client,
	}, nil
}
