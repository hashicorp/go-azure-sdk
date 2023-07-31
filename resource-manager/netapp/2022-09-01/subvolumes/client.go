package subvolumes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubVolumesClient struct {
	Client *resourcemanager.Client
}

func NewSubVolumesClientWithBaseURI(api sdkEnv.Api) (*SubVolumesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "subvolumes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SubVolumesClient: %+v", err)
	}

	return &SubVolumesClient{
		Client: client,
	}, nil
}
