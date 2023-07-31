package customlocations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomLocationsClient struct {
	Client *resourcemanager.Client
}

func NewCustomLocationsClientWithBaseURI(api sdkEnv.Api) (*CustomLocationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "customlocations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomLocationsClient: %+v", err)
	}

	return &CustomLocationsClient{
		Client: client,
	}, nil
}
