// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/availableskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/bandwidthschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/containers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/devices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/monitoringconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/nodes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/orders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/shares"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/storageaccountcredentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/storageaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/triggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2020-12-01/users"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Addons                    *addons.AddonsClient
	Alerts                    *alerts.AlertsClient
	AvailableSkus             *availableskus.AvailableSkusClient
	BandwidthSchedules        *bandwidthschedules.BandwidthSchedulesClient
	Containers                *containers.ContainersClient
	Devices                   *devices.DevicesClient
	Jobs                      *jobs.JobsClient
	MonitoringConfig          *monitoringconfig.MonitoringConfigClient
	Nodes                     *nodes.NodesClient
	Orders                    *orders.OrdersClient
	Roles                     *roles.RolesClient
	Shares                    *shares.SharesClient
	StorageAccountCredentials *storageaccountcredentials.StorageAccountCredentialsClient
	StorageAccounts           *storageaccounts.StorageAccountsClient
	Triggers                  *triggers.TriggersClient
	Users                     *users.UsersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	addonsClient := addons.NewAddonsClientWithBaseURI(endpoint)
	configureAuthFunc(&addonsClient.Client)

	alertsClient := alerts.NewAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsClient.Client)

	availableSkusClient := availableskus.NewAvailableSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&availableSkusClient.Client)

	bandwidthSchedulesClient := bandwidthschedules.NewBandwidthSchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&bandwidthSchedulesClient.Client)

	containersClient := containers.NewContainersClientWithBaseURI(endpoint)
	configureAuthFunc(&containersClient.Client)

	devicesClient := devices.NewDevicesClientWithBaseURI(endpoint)
	configureAuthFunc(&devicesClient.Client)

	jobsClient := jobs.NewJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobsClient.Client)

	monitoringConfigClient := monitoringconfig.NewMonitoringConfigClientWithBaseURI(endpoint)
	configureAuthFunc(&monitoringConfigClient.Client)

	nodesClient := nodes.NewNodesClientWithBaseURI(endpoint)
	configureAuthFunc(&nodesClient.Client)

	ordersClient := orders.NewOrdersClientWithBaseURI(endpoint)
	configureAuthFunc(&ordersClient.Client)

	rolesClient := roles.NewRolesClientWithBaseURI(endpoint)
	configureAuthFunc(&rolesClient.Client)

	sharesClient := shares.NewSharesClientWithBaseURI(endpoint)
	configureAuthFunc(&sharesClient.Client)

	storageAccountCredentialsClient := storageaccountcredentials.NewStorageAccountCredentialsClientWithBaseURI(endpoint)
	configureAuthFunc(&storageAccountCredentialsClient.Client)

	storageAccountsClient := storageaccounts.NewStorageAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&storageAccountsClient.Client)

	triggersClient := triggers.NewTriggersClientWithBaseURI(endpoint)
	configureAuthFunc(&triggersClient.Client)

	usersClient := users.NewUsersClientWithBaseURI(endpoint)
	configureAuthFunc(&usersClient.Client)

	return Client{
		Addons:                    &addonsClient,
		Alerts:                    &alertsClient,
		AvailableSkus:             &availableSkusClient,
		BandwidthSchedules:        &bandwidthSchedulesClient,
		Containers:                &containersClient,
		Devices:                   &devicesClient,
		Jobs:                      &jobsClient,
		MonitoringConfig:          &monitoringConfigClient,
		Nodes:                     &nodesClient,
		Orders:                    &ordersClient,
		Roles:                     &rolesClient,
		Shares:                    &sharesClient,
		StorageAccountCredentials: &storageAccountCredentialsClient,
		StorageAccounts:           &storageAccountsClient,
		Triggers:                  &triggersClient,
		Users:                     &usersClient,
	}
}
