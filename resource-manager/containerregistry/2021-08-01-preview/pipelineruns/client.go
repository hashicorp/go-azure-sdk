package pipelineruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunsClient struct {
	Client *resourcemanager.Client
}

func NewPipelineRunsClientWithBaseURI(api sdkEnv.Api) (*PipelineRunsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "pipelineruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelineRunsClient: %+v", err)
	}

	return &PipelineRunsClient{
		Client: client,
	}, nil
}
