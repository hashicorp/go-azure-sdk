package capabilitiesbyserver

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilitiesByServerClient struct {
	Client *resourcemanager.Client
}

func NewCapabilitiesByServerClientWithBaseURI(sdkApi sdkEnv.Api) (*CapabilitiesByServerClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "capabilitiesbyserver", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CapabilitiesByServerClient: %+v", err)
	}

	return &CapabilitiesByServerClient{
		Client: client,
	}, nil
}
