// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_03_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/availableskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/bandwidthschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/containers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/devicecapacitycheck"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/devicecapacityinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/devices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/diagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/monitoringconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/nodes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/orders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/shares"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/storageaccountcredentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/storageaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/supportpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/triggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2022-03-01/users"
)

type Client struct {
	Addons                    *addons.AddonsClient
	Alerts                    *alerts.AlertsClient
	AvailableSkus             *availableskus.AvailableSkusClient
	BandwidthSchedules        *bandwidthschedules.BandwidthSchedulesClient
	Containers                *containers.ContainersClient
	DeviceCapacityCheck       *devicecapacitycheck.DeviceCapacityCheckClient
	DeviceCapacityInfo        *devicecapacityinfo.DeviceCapacityInfoClient
	Devices                   *devices.DevicesClient
	DiagnosticSettings        *diagnosticsettings.DiagnosticSettingsClient
	Jobs                      *jobs.JobsClient
	MonitoringConfig          *monitoringconfig.MonitoringConfigClient
	Nodes                     *nodes.NodesClient
	Orders                    *orders.OrdersClient
	Roles                     *roles.RolesClient
	Shares                    *shares.SharesClient
	StorageAccountCredentials *storageaccountcredentials.StorageAccountCredentialsClient
	StorageAccounts           *storageaccounts.StorageAccountsClient
	SupportPackages           *supportpackages.SupportPackagesClient
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

	deviceCapacityCheckClient := devicecapacitycheck.NewDeviceCapacityCheckClientWithBaseURI(endpoint)
	configureAuthFunc(&deviceCapacityCheckClient.Client)

	deviceCapacityInfoClient := devicecapacityinfo.NewDeviceCapacityInfoClientWithBaseURI(endpoint)
	configureAuthFunc(&deviceCapacityInfoClient.Client)

	devicesClient := devices.NewDevicesClientWithBaseURI(endpoint)
	configureAuthFunc(&devicesClient.Client)

	diagnosticSettingsClient := diagnosticsettings.NewDiagnosticSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticSettingsClient.Client)

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

	supportPackagesClient := supportpackages.NewSupportPackagesClientWithBaseURI(endpoint)
	configureAuthFunc(&supportPackagesClient.Client)

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
		DeviceCapacityCheck:       &deviceCapacityCheckClient,
		DeviceCapacityInfo:        &deviceCapacityInfoClient,
		Devices:                   &devicesClient,
		DiagnosticSettings:        &diagnosticSettingsClient,
		Jobs:                      &jobsClient,
		MonitoringConfig:          &monitoringConfigClient,
		Nodes:                     &nodesClient,
		Orders:                    &ordersClient,
		Roles:                     &rolesClient,
		Shares:                    &sharesClient,
		StorageAccountCredentials: &storageAccountCredentialsClient,
		StorageAccounts:           &storageAccountsClient,
		SupportPackages:           &supportPackagesClient,
		Triggers:                  &triggersClient,
		Users:                     &usersClient,
	}
}
