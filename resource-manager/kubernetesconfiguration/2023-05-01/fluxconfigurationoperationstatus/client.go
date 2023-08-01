package fluxconfigurationoperationstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FluxConfigurationOperationStatusClient struct {
	Client *resourcemanager.Client
}

func NewFluxConfigurationOperationStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*FluxConfigurationOperationStatusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "fluxconfigurationoperationstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FluxConfigurationOperationStatusClient: %+v", err)
	}

	return &FluxConfigurationOperationStatusClient{
		Client: client,
	}, nil
}
