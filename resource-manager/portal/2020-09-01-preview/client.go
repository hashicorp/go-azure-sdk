package v2020_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/dashboards"
	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/listtenantconfigurationviolationsoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/tenantconfigurations"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Dashboards                                  *dashboards.DashboardsClient
	ListTenantConfigurationViolationsOperations *listtenantconfigurationviolationsoperations.ListTenantConfigurationViolationsOperationsClient
	TenantConfigurations                        *tenantconfigurations.TenantConfigurationsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dashboardsClient, err := dashboards.NewDashboardsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dashboards client: %+v", err)
	}
	configureFunc(dashboardsClient.Client)

	listTenantConfigurationViolationsOperationsClient, err := listtenantconfigurationviolationsoperations.NewListTenantConfigurationViolationsOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListTenantConfigurationViolationsOperations client: %+v", err)
	}
	configureFunc(listTenantConfigurationViolationsOperationsClient.Client)

	tenantConfigurationsClient, err := tenantconfigurations.NewTenantConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TenantConfigurations client: %+v", err)
	}
	configureFunc(tenantConfigurationsClient.Client)

	return &Client{
		Dashboards: dashboardsClient,
		ListTenantConfigurationViolationsOperations: listTenantConfigurationViolationsOperationsClient,
		TenantConfigurations:                        tenantConfigurationsClient,
	}, nil
}
