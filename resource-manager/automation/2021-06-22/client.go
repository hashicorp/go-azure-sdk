package v2021_06_22

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/automationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/hybridrunbookworker"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/hybridrunbookworkergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/listkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/statistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2021-06-22/usages"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	automationAccountClient := automationaccount.NewAutomationAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&automationAccountClient.Client)

	hybridRunbookWorkerClient := hybridrunbookworker.NewHybridRunbookWorkerClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridRunbookWorkerClient.Client)

	hybridRunbookWorkerGroupClient := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridRunbookWorkerGroupClient.Client)

	listKeysClient := listkeys.NewListKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&listKeysClient.Client)

	operationsClient := operations.NewOperationsClientWithBaseURI(endpoint)
	configureAuthFunc(&operationsClient.Client)

	statisticsClient := statistics.NewStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&statisticsClient.Client)

	usagesClient := usages.NewUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&usagesClient.Client)

	return Client{
		AutomationAccount:        &automationAccountClient,
		HybridRunbookWorker:      &hybridRunbookWorkerClient,
		HybridRunbookWorkerGroup: &hybridRunbookWorkerGroupClient,
		ListKeys:                 &listKeysClient,
		Operations:               &operationsClient,
		Statistics:               &statisticsClient,
		Usages:                   &usagesClient,
	}
}
