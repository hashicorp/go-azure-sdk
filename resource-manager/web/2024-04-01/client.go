package v2024_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/appservicecertificateorders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/appserviceenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/appserviceplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/certificateordersdiagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/deletedwebapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/diagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/domains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/global"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/kubeenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/provider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/recommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/resourcehealthmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/resourceproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/staticsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/topleveldomains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/webapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/workflowrunactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/workflowruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/workflows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/workflowtriggerhistories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/workflowtriggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2024-04-01/workflowversions"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppServiceCertificateOrders  *appservicecertificateorders.AppServiceCertificateOrdersClient
	AppServiceEnvironments       *appserviceenvironments.AppServiceEnvironmentsClient
	AppServicePlans              *appserviceplans.AppServicePlansClient
	CertificateOrdersDiagnostics *certificateordersdiagnostics.CertificateOrdersDiagnosticsClient
	Certificates                 *certificates.CertificatesClient
	DeletedWebApps               *deletedwebapps.DeletedWebAppsClient
	Diagnostics                  *diagnostics.DiagnosticsClient
	Domains                      *domains.DomainsClient
	Global                       *global.GlobalClient
	KubeEnvironments             *kubeenvironments.KubeEnvironmentsClient
	Provider                     *provider.ProviderClient
	Recommendations              *recommendations.RecommendationsClient
	ResourceHealthMetadata       *resourcehealthmetadata.ResourceHealthMetadataClient
	ResourceProviders            *resourceproviders.ResourceProvidersClient
	StaticSites                  *staticsites.StaticSitesClient
	TopLevelDomains              *topleveldomains.TopLevelDomainsClient
	WebApps                      *webapps.WebAppsClient
	WorkflowRunActions           *workflowrunactions.WorkflowRunActionsClient
	WorkflowRuns                 *workflowruns.WorkflowRunsClient
	WorkflowTriggerHistories     *workflowtriggerhistories.WorkflowTriggerHistoriesClient
	WorkflowTriggers             *workflowtriggers.WorkflowTriggersClient
	WorkflowVersions             *workflowversions.WorkflowVersionsClient
	Workflows                    *workflows.WorkflowsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appServiceCertificateOrdersClient, err := appservicecertificateorders.NewAppServiceCertificateOrdersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServiceCertificateOrders client: %+v", err)
	}
	configureFunc(appServiceCertificateOrdersClient.Client)

	appServiceEnvironmentsClient, err := appserviceenvironments.NewAppServiceEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServiceEnvironments client: %+v", err)
	}
	configureFunc(appServiceEnvironmentsClient.Client)

	appServicePlansClient, err := appserviceplans.NewAppServicePlansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppServicePlans client: %+v", err)
	}
	configureFunc(appServicePlansClient.Client)

	certificateOrdersDiagnosticsClient, err := certificateordersdiagnostics.NewCertificateOrdersDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CertificateOrdersDiagnostics client: %+v", err)
	}
	configureFunc(certificateOrdersDiagnosticsClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	deletedWebAppsClient, err := deletedwebapps.NewDeletedWebAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedWebApps client: %+v", err)
	}
	configureFunc(deletedWebAppsClient.Client)

	diagnosticsClient, err := diagnostics.NewDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Diagnostics client: %+v", err)
	}
	configureFunc(diagnosticsClient.Client)

	domainsClient, err := domains.NewDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Domains client: %+v", err)
	}
	configureFunc(domainsClient.Client)

	globalClient, err := global.NewGlobalClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Global client: %+v", err)
	}
	configureFunc(globalClient.Client)

	kubeEnvironmentsClient, err := kubeenvironments.NewKubeEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building KubeEnvironments client: %+v", err)
	}
	configureFunc(kubeEnvironmentsClient.Client)

	providerClient, err := provider.NewProviderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Provider client: %+v", err)
	}
	configureFunc(providerClient.Client)

	recommendationsClient, err := recommendations.NewRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Recommendations client: %+v", err)
	}
	configureFunc(recommendationsClient.Client)

	resourceHealthMetadataClient, err := resourcehealthmetadata.NewResourceHealthMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceHealthMetadata client: %+v", err)
	}
	configureFunc(resourceHealthMetadataClient.Client)

	resourceProvidersClient, err := resourceproviders.NewResourceProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceProviders client: %+v", err)
	}
	configureFunc(resourceProvidersClient.Client)

	staticSitesClient, err := staticsites.NewStaticSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StaticSites client: %+v", err)
	}
	configureFunc(staticSitesClient.Client)

	topLevelDomainsClient, err := topleveldomains.NewTopLevelDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TopLevelDomains client: %+v", err)
	}
	configureFunc(topLevelDomainsClient.Client)

	webAppsClient, err := webapps.NewWebAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WebApps client: %+v", err)
	}
	configureFunc(webAppsClient.Client)

	workflowRunActionsClient, err := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRunActions client: %+v", err)
	}
	configureFunc(workflowRunActionsClient.Client)

	workflowRunsClient, err := workflowruns.NewWorkflowRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowRuns client: %+v", err)
	}
	configureFunc(workflowRunsClient.Client)

	workflowTriggerHistoriesClient, err := workflowtriggerhistories.NewWorkflowTriggerHistoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowTriggerHistories client: %+v", err)
	}
	configureFunc(workflowTriggerHistoriesClient.Client)

	workflowTriggersClient, err := workflowtriggers.NewWorkflowTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowTriggers client: %+v", err)
	}
	configureFunc(workflowTriggersClient.Client)

	workflowVersionsClient, err := workflowversions.NewWorkflowVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkflowVersions client: %+v", err)
	}
	configureFunc(workflowVersionsClient.Client)

	workflowsClient, err := workflows.NewWorkflowsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workflows client: %+v", err)
	}
	configureFunc(workflowsClient.Client)

	return &Client{
		AppServiceCertificateOrders:  appServiceCertificateOrdersClient,
		AppServiceEnvironments:       appServiceEnvironmentsClient,
		AppServicePlans:              appServicePlansClient,
		CertificateOrdersDiagnostics: certificateOrdersDiagnosticsClient,
		Certificates:                 certificatesClient,
		DeletedWebApps:               deletedWebAppsClient,
		Diagnostics:                  diagnosticsClient,
		Domains:                      domainsClient,
		Global:                       globalClient,
		KubeEnvironments:             kubeEnvironmentsClient,
		Provider:                     providerClient,
		Recommendations:              recommendationsClient,
		ResourceHealthMetadata:       resourceHealthMetadataClient,
		ResourceProviders:            resourceProvidersClient,
		StaticSites:                  staticSitesClient,
		TopLevelDomains:              topLevelDomainsClient,
		WebApps:                      webAppsClient,
		WorkflowRunActions:           workflowRunActionsClient,
		WorkflowRuns:                 workflowRunsClient,
		WorkflowTriggerHistories:     workflowTriggerHistoriesClient,
		WorkflowTriggers:             workflowTriggersClient,
		WorkflowVersions:             workflowVersionsClient,
		Workflows:                    workflowsClient,
	}, nil
}
