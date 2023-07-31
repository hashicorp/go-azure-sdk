package deletedwebapps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedWebAppsClient struct {
	Client *resourcemanager.Client
}

func NewDeletedWebAppsClientWithBaseURI(api sdkEnv.Api) (*DeletedWebAppsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "deletedwebapps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedWebAppsClient: %+v", err)
	}

	return &DeletedWebAppsClient{
		Client: client,
	}, nil
}
