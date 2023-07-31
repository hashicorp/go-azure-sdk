package taskruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskRunsClient struct {
	Client *resourcemanager.Client
}

func NewTaskRunsClientWithBaseURI(api sdkEnv.Api) (*TaskRunsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "taskruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TaskRunsClient: %+v", err)
	}

	return &TaskRunsClient{
		Client: client,
	}, nil
}
