package snapshotpolicylistvolumes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnapshotPolicyListVolumesClient struct {
	Client *resourcemanager.Client
}

func NewSnapshotPolicyListVolumesClientWithBaseURI(api sdkEnv.Api) (*SnapshotPolicyListVolumesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "snapshotpolicylistvolumes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SnapshotPolicyListVolumesClient: %+v", err)
	}

	return &SnapshotPolicyListVolumesClient{
		Client: client,
	}, nil
}
