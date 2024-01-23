package v2020_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/dashboard"
	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/listtenantconfigurationviolations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2020-09-01-preview/tenantconfiguration"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Dashboard                         *dashboard.DashboardClient
	ListTenantConfigurationViolations *listtenantconfigurationviolations.ListTenantConfigurationViolationsClient
	TenantConfiguration               *tenantconfiguration.TenantConfigurationClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dashboardClient, err := dashboard.NewDashboardClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dashboard client: %+v", err)
	}
	configureFunc(dashboardClient.Client)

	listTenantConfigurationViolationsClient, err := listtenantconfigurationviolations.NewListTenantConfigurationViolationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListTenantConfigurationViolations client: %+v", err)
	}
	configureFunc(listTenantConfigurationViolationsClient.Client)

	tenantConfigurationClient, err := tenantconfiguration.NewTenantConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TenantConfiguration client: %+v", err)
	}
	configureFunc(tenantConfigurationClient.Client)

	return &Client{
		Dashboard:                         dashboardClient,
		ListTenantConfigurationViolations: listTenantConfigurationViolationsClient,
		TenantConfiguration:               tenantConfigurationClient,
	}, nil
}
