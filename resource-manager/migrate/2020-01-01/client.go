package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

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

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	hyperVClusterClient, err := hypervcluster.NewHyperVClusterClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HyperVCluster client: %+v", err)
	}
	configureFunc(hyperVClusterClient.Client)

	hyperVHostClient, err := hypervhost.NewHyperVHostClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HyperVHost client: %+v", err)
	}
	configureFunc(hyperVHostClient.Client)

	hyperVJobsClient, err := hypervjobs.NewHyperVJobsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HyperVJobs client: %+v", err)
	}
	configureFunc(hyperVJobsClient.Client)

	hyperVMachinesClient, err := hypervmachines.NewHyperVMachinesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HyperVMachines client: %+v", err)
	}
	configureFunc(hyperVMachinesClient.Client)

	hyperVRunAsAccountsClient, err := hypervrunasaccounts.NewHyperVRunAsAccountsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HyperVRunAsAccounts client: %+v", err)
	}
	configureFunc(hyperVRunAsAccountsClient.Client)

	hyperVSitesClient, err := hypervsites.NewHyperVSitesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HyperVSites client: %+v", err)
	}
	configureFunc(hyperVSitesClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	machinesClient, err := machines.NewMachinesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Machines client: %+v", err)
	}
	configureFunc(machinesClient.Client)

	runAsAccountsClient, err := runasaccounts.NewRunAsAccountsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RunAsAccounts client: %+v", err)
	}
	configureFunc(runAsAccountsClient.Client)

	sitesClient, err := sites.NewSitesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Sites client: %+v", err)
	}
	configureFunc(sitesClient.Client)

	vCenterClient, err := vcenter.NewVCenterClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VCenter client: %+v", err)
	}
	configureFunc(vCenterClient.Client)

	return &Client{
		HyperVCluster:       hyperVClusterClient,
		HyperVHost:          hyperVHostClient,
		HyperVJobs:          hyperVJobsClient,
		HyperVMachines:      hyperVMachinesClient,
		HyperVRunAsAccounts: hyperVRunAsAccountsClient,
		HyperVSites:         hyperVSitesClient,
		Jobs:                jobsClient,
		Machines:            machinesClient,
		RunAsAccounts:       runAsAccountsClient,
		Sites:               sitesClient,
		VCenter:             vCenterClient,
	}, nil
}
