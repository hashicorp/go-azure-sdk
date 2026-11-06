package containerappslabelhistory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppsLabelHistoryClient struct {
	Client *resourcemanager.Client
}

func NewContainerAppsLabelHistoryClientWithBaseURI(sdkApi sdkEnv.Api) (*ContainerAppsLabelHistoryClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "containerappslabelhistory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContainerAppsLabelHistoryClient: %+v", err)
	}

	return &ContainerAppsLabelHistoryClient{
		Client: client,
	}, nil
}
