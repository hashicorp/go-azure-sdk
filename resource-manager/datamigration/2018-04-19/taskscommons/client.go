package taskscommons

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TasksCommonsClient struct {
	Client *resourcemanager.Client
}

func NewTasksCommonsClientWithBaseURI(sdkApi sdkEnv.Api) (*TasksCommonsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "taskscommons", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TasksCommonsClient: %+v", err)
	}

	return &TasksCommonsClient{
		Client: client,
	}, nil
}
