package workloadgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadGroupsClient struct {
	Client *resourcemanager.Client
}

func NewWorkloadGroupsClientWithBaseURI(api sdkEnv.Api) (*WorkloadGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "workloadgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkloadGroupsClient: %+v", err)
	}

	return &WorkloadGroupsClient{
		Client: client,
	}, nil
}
