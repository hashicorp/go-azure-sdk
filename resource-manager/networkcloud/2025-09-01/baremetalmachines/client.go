package baremetalmachines

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachinesClient struct {
	Client *resourcemanager.Client
}

func NewBareMetalMachinesClientWithBaseURI(sdkApi sdkEnv.Api) (*BareMetalMachinesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "baremetalmachines", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BareMetalMachinesClient: %+v", err)
	}

	return &BareMetalMachinesClient{
		Client: client,
	}, nil
}
