package managedinstanceoperations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceOperationsClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstanceOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedInstanceOperationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "managedinstanceoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstanceOperationsClient: %+v", err)
	}

	return &ManagedInstanceOperationsClient{
		Client: client,
	}, nil
}
