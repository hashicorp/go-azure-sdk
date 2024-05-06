package volumesrelocation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VolumesRelocationClient struct {
	Client *resourcemanager.Client
}

func NewVolumesRelocationClientWithBaseURI(sdkApi sdkEnv.Api) (*VolumesRelocationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "volumesrelocation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VolumesRelocationClient: %+v", err)
	}

	return &VolumesRelocationClient{
		Client: client,
	}, nil
}
