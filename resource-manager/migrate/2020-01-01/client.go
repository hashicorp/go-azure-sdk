// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervcluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervjobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervrunasaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/runasaccounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/sites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/migrate/2020-01-01/vcenter"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	HyperVCluster       *hypervcluster.HyperVClusterClient
	HyperVHost          *hypervhost.HyperVHostClient
	HyperVJobs          *hypervjobs.HyperVJobsClient
	HyperVMachines      *hypervmachines.HyperVMachinesClient
	HyperVRunAsAccounts *hypervrunasaccounts.HyperVRunAsAccountsClient
	HyperVSites         *hypervsites.HyperVSitesClient
	Jobs                *jobs.JobsClient
	Machines            *machines.MachinesClient
	RunAsAccounts       *runasaccounts.RunAsAccountsClient
	Sites               *sites.SitesClient
	VCenter             *vcenter.VCenterClient
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

	runAsAccountsClient := runasaccounts.NewRunAsAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&runAsAccountsClient.Client)

	sitesClient := sites.NewSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&sitesClient.Client)

	vCenterClient := vcenter.NewVCenterClientWithBaseURI(endpoint)
	configureAuthFunc(&vCenterClient.Client)

	return Client{
		HyperVCluster:       &hyperVClusterClient,
		HyperVHost:          &hyperVHostClient,
		HyperVJobs:          &hyperVJobsClient,
		HyperVMachines:      &hyperVMachinesClient,
		HyperVRunAsAccounts: &hyperVRunAsAccountsClient,
		HyperVSites:         &hyperVSitesClient,
		Jobs:                &jobsClient,
		Machines:            &machinesClient,
		RunAsAccounts:       &runAsAccountsClient,
		Sites:               &sitesClient,
		VCenter:             &vCenterClient,
	}
}
