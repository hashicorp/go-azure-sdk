package autoscaleapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoscaleAPIsClient struct {
	Client *resourcemanager.Client
}

func NewAutoscaleAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*AutoscaleAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "autoscaleapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AutoscaleAPIsClient: %+v", err)
	}

	return &AutoscaleAPIsClient{
		Client: client,
	}, nil
}
