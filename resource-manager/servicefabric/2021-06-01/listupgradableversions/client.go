package listupgradableversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUpgradableVersionsClient struct {
	Client *resourcemanager.Client
}

func NewListUpgradableVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ListUpgradableVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "listupgradableversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ListUpgradableVersionsClient: %+v", err)
	}

	return &ListUpgradableVersionsClient{
		Client: client,
	}, nil
}
