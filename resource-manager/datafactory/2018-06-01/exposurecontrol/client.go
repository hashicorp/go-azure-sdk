package exposurecontrol

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExposureControlClient struct {
	Client *resourcemanager.Client
}

func NewExposureControlClientWithBaseURI(api sdkEnv.Api) (*ExposureControlClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "exposurecontrol", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExposureControlClient: %+v", err)
	}

	return &ExposureControlClient{
		Client: client,
	}, nil
}
