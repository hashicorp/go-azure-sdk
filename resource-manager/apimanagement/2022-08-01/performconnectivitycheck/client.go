package performconnectivitycheck

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PerformConnectivityCheckClient struct {
	Client *resourcemanager.Client
}

func NewPerformConnectivityCheckClientWithBaseURI(api sdkEnv.Api) (*PerformConnectivityCheckClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "performconnectivitycheck", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PerformConnectivityCheckClient: %+v", err)
	}

	return &PerformConnectivityCheckClient{
		Client: client,
	}, nil
}
