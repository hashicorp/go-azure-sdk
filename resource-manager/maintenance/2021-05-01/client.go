// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/applyupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/applyupdates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/configurationassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/publicmaintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maintenance/2021-05-01/updates"
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
