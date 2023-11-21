package smartgroups

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SmartGroupsClient struct {
	Client *resourcemanager.Client
}

func NewSmartGroupsClientWithBaseURI(sdkApi sdkEnv.Api) (*SmartGroupsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "smartgroups", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SmartGroupsClient: %+v", err)
	}

	return &SmartGroupsClient{
		Client: client,
	}, nil
}
