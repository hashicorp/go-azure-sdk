package autoscalevcores

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoScaleVCoresClient struct {
	Client *resourcemanager.Client
}

func NewAutoScaleVCoresClientWithBaseURI(api sdkEnv.Api) (*AutoScaleVCoresClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "autoscalevcores", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AutoScaleVCoresClient: %+v", err)
	}

	return &AutoScaleVCoresClient{
		Client: client,
	}, nil
}
