package v2019_01_01_preview

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2019-01-01-preview/dashboard"
	"github.com/hashicorp/go-azure-sdk/resource-manager/portal/2019-01-01-preview/tenantconfiguration"
)

type Client struct {
	Dashboard           *dashboard.DashboardClient
	TenantConfiguration *tenantconfiguration.TenantConfigurationClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dashboardClient := dashboard.NewDashboardClientWithBaseURI(endpoint)
	configureAuthFunc(&dashboardClient.Client)

	tenantConfigurationClient := tenantconfiguration.NewTenantConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&tenantConfigurationClient.Client)

	return Client{
		Dashboard:           &dashboardClient,
		TenantConfiguration: &tenantConfigurationClient,
	}
}
