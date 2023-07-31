package manageddatabasesensitivitylabels

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseSensitivityLabelsClient struct {
	Client *resourcemanager.Client
}

func NewManagedDatabaseSensitivityLabelsClientWithBaseURI(api sdkEnv.Api) (*ManagedDatabaseSensitivityLabelsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "manageddatabasesensitivitylabels", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedDatabaseSensitivityLabelsClient: %+v", err)
	}

	return &ManagedDatabaseSensitivityLabelsClient{
		Client: client,
	}, nil
}
