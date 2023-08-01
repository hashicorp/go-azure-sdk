package v2023_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/availableskus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/bandwidthschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/containers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/devicecapacitycheck"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/devicecapacityinfo"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/devices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/diagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/monitoringconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/nodes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/orders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/shares"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/storageaccountcredentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/storageaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/supportpackages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/triggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-07-01/users"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	addonsClient, err := addons.NewAddonsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Addons client: %+v", err)
	}
	configureFunc(addonsClient.Client)

	alertsClient, err := alerts.NewAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	availableSkusClient, err := availableskus.NewAvailableSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableSkus client: %+v", err)
	}
	configureFunc(availableSkusClient.Client)

	bandwidthSchedulesClient, err := bandwidthschedules.NewBandwidthSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BandwidthSchedules client: %+v", err)
	}
	configureFunc(bandwidthSchedulesClient.Client)

	containersClient, err := containers.NewContainersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Containers client: %+v", err)
	}
	configureFunc(containersClient.Client)

	deviceCapacityCheckClient, err := devicecapacitycheck.NewDeviceCapacityCheckClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCapacityCheck client: %+v", err)
	}
	configureFunc(deviceCapacityCheckClient.Client)

	deviceCapacityInfoClient, err := devicecapacityinfo.NewDeviceCapacityInfoClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCapacityInfo client: %+v", err)
	}
	configureFunc(deviceCapacityInfoClient.Client)

	devicesClient, err := devices.NewDevicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Devices client: %+v", err)
	}
	configureFunc(devicesClient.Client)

	diagnosticSettingsClient, err := diagnosticsettings.NewDiagnosticSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticSettings client: %+v", err)
	}
	configureFunc(diagnosticSettingsClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	monitoringConfigClient, err := monitoringconfig.NewMonitoringConfigClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoringConfig client: %+v", err)
	}
	configureFunc(monitoringConfigClient.Client)

	nodesClient, err := nodes.NewNodesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Nodes client: %+v", err)
	}
	configureFunc(nodesClient.Client)

	ordersClient, err := orders.NewOrdersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Orders client: %+v", err)
	}
	configureFunc(ordersClient.Client)

	rolesClient, err := roles.NewRolesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Roles client: %+v", err)
	}
	configureFunc(rolesClient.Client)

	sharesClient, err := shares.NewSharesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Shares client: %+v", err)
	}
	configureFunc(sharesClient.Client)

	storageAccountCredentialsClient, err := storageaccountcredentials.NewStorageAccountCredentialsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageAccountCredentials client: %+v", err)
	}
	configureFunc(storageAccountCredentialsClient.Client)

	storageAccountsClient, err := storageaccounts.NewStorageAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageAccounts client: %+v", err)
	}
	configureFunc(storageAccountsClient.Client)

	supportPackagesClient, err := supportpackages.NewSupportPackagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SupportPackages client: %+v", err)
	}
	configureFunc(supportPackagesClient.Client)

	triggersClient, err := triggers.NewTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Triggers client: %+v", err)
	}
	configureFunc(triggersClient.Client)

	usersClient, err := users.NewUsersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Users client: %+v", err)
	}
	configureFunc(usersClient.Client)

	return &Client{
		Addons:                    addonsClient,
		Alerts:                    alertsClient,
		AvailableSkus:             availableSkusClient,
		BandwidthSchedules:        bandwidthSchedulesClient,
		Containers:                containersClient,
		DeviceCapacityCheck:       deviceCapacityCheckClient,
		DeviceCapacityInfo:        deviceCapacityInfoClient,
		Devices:                   devicesClient,
		DiagnosticSettings:        diagnosticSettingsClient,
		Jobs:                      jobsClient,
		MonitoringConfig:          monitoringConfigClient,
		Nodes:                     nodesClient,
		Orders:                    ordersClient,
		Roles:                     rolesClient,
		Shares:                    sharesClient,
		StorageAccountCredentials: storageAccountCredentialsClient,
		StorageAccounts:           storageAccountsClient,
		SupportPackages:           supportPackagesClient,
		Triggers:                  triggersClient,
		Users:                     usersClient,
	}, nil
}
