package recoverypoint

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointClient struct {
	Client *resourcemanager.Client
}

func NewRecoveryPointClientWithBaseURI(api environments.Api) (*RecoveryPointClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "recoverypoint", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecoveryPointClient: %+v", err)
	}

	return &RecoveryPointClient{
		Client: client,
	}, nil
}
