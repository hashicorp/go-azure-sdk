// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_07_07

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervcluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervrunasaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/hypervsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/mastersites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/migrates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/privateendpointconnection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/runasaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/sites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-07-07/vcenter"
)

type Client struct {
	HyperVCluster             *hypervcluster.HyperVClusterClient
	HyperVHost                *hypervhost.HyperVHostClient
	HyperVJobs                *hypervjobs.HyperVJobsClient
	HyperVMachines            *hypervmachines.HyperVMachinesClient
	HyperVRunAsAccounts       *hypervrunasaccounts.HyperVRunAsAccountsClient
	HyperVSites               *hypervsites.HyperVSitesClient
	Jobs                      *jobs.JobsClient
	Machines                  *machines.MachinesClient
	MasterSites               *mastersites.MasterSitesClient
	Migrates                  *migrates.MigratesClient
	PrivateEndpointConnection *privateendpointconnection.PrivateEndpointConnectionClient
	PrivateLinkResources      *privatelinkresources.PrivateLinkResourcesClient
	RunAsAccounts             *runasaccounts.RunAsAccountsClient
	Sites                     *sites.SitesClient
	VCenter                   *vcenter.VCenterClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	hyperVClusterClient := hypervcluster.NewHyperVClusterClientWithBaseURI(endpoint)
	configureAuthFunc(&hyperVClusterClient.Client)

	hyperVHostClient := hypervhost.NewHyperVHostClientWithBaseURI(endpoint)
	configureAuthFunc(&hyperVHostClient.Client)

	hyperVJobsClient := hypervjobs.NewHyperVJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&hyperVJobsClient.Client)

	hyperVMachinesClient := hypervmachines.NewHyperVMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&hyperVMachinesClient.Client)

	hyperVRunAsAccountsClient := hypervrunasaccounts.NewHyperVRunAsAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&hyperVRunAsAccountsClient.Client)

	hyperVSitesClient := hypervsites.NewHyperVSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&hyperVSitesClient.Client)

	jobsClient := jobs.NewJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobsClient.Client)

	machinesClient := machines.NewMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&machinesClient.Client)

	masterSitesClient := mastersites.NewMasterSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&masterSitesClient.Client)

	migratesClient := migrates.NewMigratesClientWithBaseURI(endpoint)
	configureAuthFunc(&migratesClient.Client)

	privateEndpointConnectionClient := privateendpointconnection.NewPrivateEndpointConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	runAsAccountsClient := runasaccounts.NewRunAsAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&runAsAccountsClient.Client)

	sitesClient := sites.NewSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&sitesClient.Client)

	vCenterClient := vcenter.NewVCenterClientWithBaseURI(endpoint)
	configureAuthFunc(&vCenterClient.Client)

	return Client{
		HyperVCluster:             &hyperVClusterClient,
		HyperVHost:                &hyperVHostClient,
		HyperVJobs:                &hyperVJobsClient,
		HyperVMachines:            &hyperVMachinesClient,
		HyperVRunAsAccounts:       &hyperVRunAsAccountsClient,
		HyperVSites:               &hyperVSitesClient,
		Jobs:                      &jobsClient,
		Machines:                  &machinesClient,
		MasterSites:               &masterSitesClient,
		Migrates:                  &migratesClient,
		PrivateEndpointConnection: &privateEndpointConnectionClient,
		PrivateLinkResources:      &privateLinkResourcesClient,
		RunAsAccounts:             &runAsAccountsClient,
		Sites:                     &sitesClient,
		VCenter:                   &vCenterClient,
	}
}
