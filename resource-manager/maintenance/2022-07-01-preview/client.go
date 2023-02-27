// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2022_07_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/applyupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/applyupdates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/configurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/publicmaintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2022-07-01-preview/updates"
)

type Client struct {
	ApplyUpdate                     *applyupdate.ApplyUpdateClient
	ApplyUpdates                    *applyupdates.ApplyUpdatesClient
	ConfigurationAssignments        *configurationassignments.ConfigurationAssignmentsClient
	MaintenanceConfigurations       *maintenanceconfigurations.MaintenanceConfigurationsClient
	PublicMaintenanceConfigurations *publicmaintenanceconfigurations.PublicMaintenanceConfigurationsClient
	Updates                         *updates.UpdatesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applyUpdateClient := applyupdate.NewApplyUpdateClientWithBaseURI(endpoint)
	configureAuthFunc(&applyUpdateClient.Client)

	applyUpdatesClient := applyupdates.NewApplyUpdatesClientWithBaseURI(endpoint)
	configureAuthFunc(&applyUpdatesClient.Client)

	configurationAssignmentsClient := configurationassignments.NewConfigurationAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationAssignmentsClient.Client)

	maintenanceConfigurationsClient := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&maintenanceConfigurationsClient.Client)

	publicMaintenanceConfigurationsClient := publicmaintenanceconfigurations.NewPublicMaintenanceConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&publicMaintenanceConfigurationsClient.Client)

	updatesClient := updates.NewUpdatesClientWithBaseURI(endpoint)
	configureAuthFunc(&updatesClient.Client)

	return Client{
		ApplyUpdate:                     &applyUpdateClient,
		ApplyUpdates:                    &applyUpdatesClient,
		ConfigurationAssignments:        &configurationAssignmentsClient,
		MaintenanceConfigurations:       &maintenanceConfigurationsClient,
		PublicMaintenanceConfigurations: &publicMaintenanceConfigurationsClient,
		Updates:                         &updatesClient,
	}
}
