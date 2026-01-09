package v2025_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/durabletask/2025-11-01/retentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/durabletask/2025-11-01/schedulers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/durabletask/2025-11-01/taskhubs"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	RetentionPolicies *retentionpolicies.RetentionPoliciesClient
	Schedulers        *schedulers.SchedulersClient
	TaskHubs          *taskhubs.TaskHubsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	retentionPoliciesClient, err := retentionpolicies.NewRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RetentionPolicies client: %+v", err)
	}
	configureFunc(retentionPoliciesClient.Client)

	schedulersClient, err := schedulers.NewSchedulersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Schedulers client: %+v", err)
	}
	configureFunc(schedulersClient.Client)

	taskHubsClient, err := taskhubs.NewTaskHubsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TaskHubs client: %+v", err)
	}
	configureFunc(taskHubsClient.Client)

	return &Client{
		RetentionPolicies: retentionPoliciesClient,
		Schedulers:        schedulersClient,
		TaskHubs:          taskHubsClient,
	}, nil
}
