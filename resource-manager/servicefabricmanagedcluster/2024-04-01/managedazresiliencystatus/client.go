package managedazresiliencystatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAzResiliencyStatusClient struct {
	Client *resourcemanager.Client
}

func NewManagedAzResiliencyStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedAzResiliencyStatusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managedazresiliencystatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedAzResiliencyStatusClient: %+v", err)
	}

	return &ManagedAzResiliencyStatusClient{
		Client: client,
	}, nil
}
