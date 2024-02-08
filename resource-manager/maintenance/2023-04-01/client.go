package v2023_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/applyupdates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/configurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/publicmaintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2023-04-01/updates"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApplyUpdates                    *applyupdates.ApplyUpdatesClient
	ConfigurationAssignments        *configurationassignments.ConfigurationAssignmentsClient
	MaintenanceConfigurations       *maintenanceconfigurations.MaintenanceConfigurationsClient
	PublicMaintenanceConfigurations *publicmaintenanceconfigurations.PublicMaintenanceConfigurationsClient
	Updates                         *updates.UpdatesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applyUpdatesClient, err := applyupdates.NewApplyUpdatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplyUpdates client: %+v", err)
	}
	configureFunc(applyUpdatesClient.Client)

	configurationAssignmentsClient, err := configurationassignments.NewConfigurationAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationAssignments client: %+v", err)
	}
	configureFunc(configurationAssignmentsClient.Client)

	maintenanceConfigurationsClient, err := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceConfigurations client: %+v", err)
	}
	configureFunc(maintenanceConfigurationsClient.Client)

	publicMaintenanceConfigurationsClient, err := publicmaintenanceconfigurations.NewPublicMaintenanceConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PublicMaintenanceConfigurations client: %+v", err)
	}
	configureFunc(publicMaintenanceConfigurationsClient.Client)

	updatesClient, err := updates.NewUpdatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Updates client: %+v", err)
	}
	configureFunc(updatesClient.Client)

	return &Client{
		ApplyUpdates:                    applyUpdatesClient,
		ConfigurationAssignments:        configurationAssignmentsClient,
		MaintenanceConfigurations:       maintenanceConfigurationsClient,
		PublicMaintenanceConfigurations: publicMaintenanceConfigurationsClient,
		Updates:                         updatesClient,
	}, nil
}
