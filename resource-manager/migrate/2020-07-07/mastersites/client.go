package mastersites

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MasterSitesClient struct {
	Client *resourcemanager.Client
}

func NewMasterSitesClientWithBaseURI(sdkApi sdkEnv.Api) (*MasterSitesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "mastersites", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MasterSitesClient: %+v", err)
	}

	return &MasterSitesClient{
		Client: client,
	}, nil
}
