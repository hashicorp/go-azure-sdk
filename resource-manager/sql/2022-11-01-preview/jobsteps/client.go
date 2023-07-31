package jobsteps

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobStepsClient struct {
	Client *resourcemanager.Client
}

func NewJobStepsClientWithBaseURI(api sdkEnv.Api) (*JobStepsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jobsteps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobStepsClient: %+v", err)
	}

	return &JobStepsClient{
		Client: client,
	}, nil
}
