package hierarchysettingsoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HierarchySettingsOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewHierarchySettingsOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*HierarchySettingsOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hierarchysettingsoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HierarchySettingsOperationGroupClient: %+v", err)
	}

	return &HierarchySettingsOperationGroupClient{
		Client: client,
	}, nil
}
