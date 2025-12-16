package v2023_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/addons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/bandwidthschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/containers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/databoxedgedevices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/databoxedges"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/devicecapacityinfos"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/devices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/diagnosticremotesupportsettingsoperationgroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/diagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/monitoringmetricconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/orders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/shares"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/storageaccountcredentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/storageaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/triggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/updatesummaries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databoxedge/2023-12-01/users"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Addons                                        *addons.AddonsClient
	Alerts                                        *alerts.AlertsClient
	BandwidthSchedules                            *bandwidthschedules.BandwidthSchedulesClient
	Containers                                    *containers.ContainersClient
	DataBoxEdgeDevices                            *databoxedgedevices.DataBoxEdgeDevicesClient
	Databoxedges                                  *databoxedges.DataboxedgesClient
	DeviceCapacityInfos                           *devicecapacityinfos.DeviceCapacityInfosClient
	Devices                                       *devices.DevicesClient
	DiagnosticRemoteSupportSettingsOperationGroup *diagnosticremotesupportsettingsoperationgroup.DiagnosticRemoteSupportSettingsOperationGroupClient
	DiagnosticSettings                            *diagnosticsettings.DiagnosticSettingsClient
	Jobs                                          *jobs.JobsClient
	MonitoringMetricConfigurations                *monitoringmetricconfigurations.MonitoringMetricConfigurationsClient
	Orders                                        *orders.OrdersClient
	Roles                                         *roles.RolesClient
	Shares                                        *shares.SharesClient
	StorageAccountCredentials                     *storageaccountcredentials.StorageAccountCredentialsClient
	StorageAccounts                               *storageaccounts.StorageAccountsClient
	Triggers                                      *triggers.TriggersClient
	UpdateSummaries                               *updatesummaries.UpdateSummariesClient
	Users                                         *users.UsersClient
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

	dataBoxEdgeDevicesClient, err := databoxedgedevices.NewDataBoxEdgeDevicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataBoxEdgeDevices client: %+v", err)
	}
	configureFunc(dataBoxEdgeDevicesClient.Client)

	databoxedgesClient, err := databoxedges.NewDataboxedgesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Databoxedges client: %+v", err)
	}
	configureFunc(databoxedgesClient.Client)

	deviceCapacityInfosClient, err := devicecapacityinfos.NewDeviceCapacityInfosClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceCapacityInfos client: %+v", err)
	}
	configureFunc(deviceCapacityInfosClient.Client)

	devicesClient, err := devices.NewDevicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Devices client: %+v", err)
	}
	configureFunc(devicesClient.Client)

	diagnosticRemoteSupportSettingsOperationGroupClient, err := diagnosticremotesupportsettingsoperationgroup.NewDiagnosticRemoteSupportSettingsOperationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticRemoteSupportSettingsOperationGroup client: %+v", err)
	}
	configureFunc(diagnosticRemoteSupportSettingsOperationGroupClient.Client)

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

	monitoringMetricConfigurationsClient, err := monitoringmetricconfigurations.NewMonitoringMetricConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MonitoringMetricConfigurations client: %+v", err)
	}
	configureFunc(monitoringMetricConfigurationsClient.Client)

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

	triggersClient, err := triggers.NewTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Triggers client: %+v", err)
	}
	configureFunc(triggersClient.Client)

	updateSummariesClient, err := updatesummaries.NewUpdateSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateSummaries client: %+v", err)
	}
	configureFunc(updateSummariesClient.Client)

	usersClient, err := users.NewUsersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Users client: %+v", err)
	}
	configureFunc(usersClient.Client)

	return &Client{
		Addons:              addonsClient,
		Alerts:              alertsClient,
		BandwidthSchedules:  bandwidthSchedulesClient,
		Containers:          containersClient,
		DataBoxEdgeDevices:  dataBoxEdgeDevicesClient,
		Databoxedges:        databoxedgesClient,
		DeviceCapacityInfos: deviceCapacityInfosClient,
		Devices:             devicesClient,
		DiagnosticRemoteSupportSettingsOperationGroup: diagnosticRemoteSupportSettingsOperationGroupClient,
		DiagnosticSettings:                            diagnosticSettingsClient,
		Jobs:                                          jobsClient,
		MonitoringMetricConfigurations:                monitoringMetricConfigurationsClient,
		Orders:                                        ordersClient,
		Roles:                                         rolesClient,
		Shares:                                        sharesClient,
		StorageAccountCredentials:                     storageAccountCredentialsClient,
		StorageAccounts:                               storageAccountsClient,
		Triggers:                                      triggersClient,
		UpdateSummaries:                               updateSummariesClient,
		Users:                                         usersClient,
	}, nil
}
