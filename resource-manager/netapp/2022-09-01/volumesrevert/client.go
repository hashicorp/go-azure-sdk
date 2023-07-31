package volumesrevert

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumesRevertClient struct {
	Client *resourcemanager.Client
}

func NewVolumesRevertClientWithBaseURI(api sdkEnv.Api) (*VolumesRevertClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "volumesrevert", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VolumesRevertClient: %+v", err)
	}

	return &VolumesRevertClient{
		Client: client,
	}, nil
}
