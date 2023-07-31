package checknameavailabilitynamespaces

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckNameAvailabilityNamespacesClient struct {
	Client *resourcemanager.Client
}

func NewCheckNameAvailabilityNamespacesClientWithBaseURI(api sdkEnv.Api) (*CheckNameAvailabilityNamespacesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "checknameavailabilitynamespaces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CheckNameAvailabilityNamespacesClient: %+v", err)
	}

	return &CheckNameAvailabilityNamespacesClient{
		Client: client,
	}, nil
}
