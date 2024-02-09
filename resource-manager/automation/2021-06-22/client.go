package v2021_06_22

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/automationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/hybridrunbookworker"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/hybridrunbookworkergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/listkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/statistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AutomationAccount        *automationaccount.AutomationAccountClient
	HybridRunbookWorker      *hybridrunbookworker.HybridRunbookWorkerClient
	HybridRunbookWorkerGroup *hybridrunbookworkergroup.HybridRunbookWorkerGroupClient
	ListKeys                 *listkeys.ListKeysClient
	Operations               *operations.OperationsClient
	Statistics               *statistics.StatisticsClient
	Usages                   *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	automationAccountClient, err := automationaccount.NewAutomationAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutomationAccount client: %+v", err)
	}
	configureFunc(automationAccountClient.Client)

	hybridRunbookWorkerClient, err := hybridrunbookworker.NewHybridRunbookWorkerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridRunbookWorker client: %+v", err)
	}
	configureFunc(hybridRunbookWorkerClient.Client)

	hybridRunbookWorkerGroupClient, err := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridRunbookWorkerGroup client: %+v", err)
	}
	configureFunc(hybridRunbookWorkerGroupClient.Client)

	listKeysClient, err := listkeys.NewListKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListKeys client: %+v", err)
	}
	configureFunc(listKeysClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	statisticsClient, err := statistics.NewStatisticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Statistics client: %+v", err)
	}
	configureFunc(statisticsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		AutomationAccount:        automationAccountClient,
		HybridRunbookWorker:      hybridRunbookWorkerClient,
		HybridRunbookWorkerGroup: hybridRunbookWorkerGroupClient,
		ListKeys:                 listKeysClient,
		Operations:               operationsClient,
		Statistics:               statisticsClient,
		Usages:                   usagesClient,
	}, nil
}
