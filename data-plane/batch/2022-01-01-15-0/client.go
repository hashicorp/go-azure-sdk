package v2022_01_01_15_0

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/accounts"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/applications"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/certificates"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/computenodes"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/files"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/jobs"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/jobschedules"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/pools"
	"github.com/hashicorp/go-azure-sdk/data-plane/batch/2022-01-01-15-0/tasks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	Accounts     *accounts.AccountsClient
	Applications *applications.ApplicationsClient
	Certificates *certificates.CertificatesClient
	ComputeNodes *computenodes.ComputeNodesClient
	Files        *files.FilesClient
	JobSchedules *jobschedules.JobSchedulesClient
	Jobs         *jobs.JobsClient
	Pools        *pools.PoolsClient
	Tasks        *tasks.TasksClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	applicationsClient, err := applications.NewApplicationsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Applications client: %+v", err)
	}
	configureFunc(applicationsClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	computeNodesClient, err := computenodes.NewComputeNodesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building ComputeNodes client: %+v", err)
	}
	configureFunc(computeNodesClient.Client)

	filesClient, err := files.NewFilesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Files client: %+v", err)
	}
	configureFunc(filesClient.Client)

	jobSchedulesClient, err := jobschedules.NewJobSchedulesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building JobSchedules client: %+v", err)
	}
	configureFunc(jobSchedulesClient.Client)

	jobsClient, err := jobs.NewJobsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	poolsClient, err := pools.NewPoolsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Pools client: %+v", err)
	}
	configureFunc(poolsClient.Client)

	tasksClient, err := tasks.NewTasksClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Tasks client: %+v", err)
	}
	configureFunc(tasksClient.Client)

	return &Client{
		Accounts:     accountsClient,
		Applications: applicationsClient,
		Certificates: certificatesClient,
		ComputeNodes: computeNodesClient,
		Files:        filesClient,
		JobSchedules: jobSchedulesClient,
		Jobs:         jobsClient,
		Pools:        poolsClient,
		Tasks:        tasksClient,
	}, nil
}
