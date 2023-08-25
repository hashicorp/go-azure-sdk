package managementconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewManagementConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagementConfigurationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managementconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagementConfigurationClient: %+v", err)
	}

	return &ManagementConfigurationClient{
		Client: client,
	}, nil
}
