package v2022_07_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/applyupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/applyupdates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/configurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/publicmaintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/updates"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ApplyUpdate                     *applyupdate.ApplyUpdateClient
	ApplyUpdates                    *applyupdates.ApplyUpdatesClient
	ConfigurationAssignments        *configurationassignments.ConfigurationAssignmentsClient
	MaintenanceConfigurations       *maintenanceconfigurations.MaintenanceConfigurationsClient
	PublicMaintenanceConfigurations *publicmaintenanceconfigurations.PublicMaintenanceConfigurationsClient
	Updates                         *updates.UpdatesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applyUpdateClient, err := applyupdate.NewApplyUpdateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApplyUpdate client: %+v", err)
	}
	configureFunc(applyUpdateClient.Client)

	applyUpdatesClient, err := applyupdates.NewApplyUpdatesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ApplyUpdates client: %+v", err)
	}
	configureFunc(applyUpdatesClient.Client)

	configurationAssignmentsClient, err := configurationassignments.NewConfigurationAssignmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationAssignments client: %+v", err)
	}
	configureFunc(configurationAssignmentsClient.Client)

	maintenanceConfigurationsClient, err := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceConfigurations client: %+v", err)
	}
	configureFunc(maintenanceConfigurationsClient.Client)

	publicMaintenanceConfigurationsClient, err := publicmaintenanceconfigurations.NewPublicMaintenanceConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PublicMaintenanceConfigurations client: %+v", err)
	}
	configureFunc(publicMaintenanceConfigurationsClient.Client)

	updatesClient, err := updates.NewUpdatesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Updates client: %+v", err)
	}
	configureFunc(updatesClient.Client)

	return &Client{
		ApplyUpdate:                     applyUpdateClient,
		ApplyUpdates:                    applyUpdatesClient,
		ConfigurationAssignments:        configurationAssignmentsClient,
		MaintenanceConfigurations:       maintenanceConfigurationsClient,
		PublicMaintenanceConfigurations: publicMaintenanceConfigurationsClient,
		Updates:                         updatesClient,
	}, nil
}
