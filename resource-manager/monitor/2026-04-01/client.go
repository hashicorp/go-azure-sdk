package v2026_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/monitor/2026-04-01/pipelinegroups"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PipelineGroups *pipelinegroups.PipelineGroupsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	pipelineGroupsClient, err := pipelinegroups.NewPipelineGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PipelineGroups client: %+v", err)
	}
	configureFunc(pipelineGroupsClient.Client)

	return &Client{
		PipelineGroups: pipelineGroupsClient,
	}, nil
}
