package v2022_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/appservicecertificateorders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/appserviceenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/appserviceplans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/certificateordersdiagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/containerapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/containerappsrevisions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/deletedwebapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/diagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/domains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/global"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/kubeenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/provider"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/recommendations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/resourcehealthmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/resourceproviders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/staticsites"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/topleveldomains"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/webapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowrunactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowtriggerhistories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowtriggers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowversions"
)

type Client struct {
	AppServiceCertificateOrders  *appservicecertificateorders.AppServiceCertificateOrdersClient
	AppServiceEnvironments       *appserviceenvironments.AppServiceEnvironmentsClient
	AppServicePlans              *appserviceplans.AppServicePlansClient
	CertificateOrdersDiagnostics *certificateordersdiagnostics.CertificateOrdersDiagnosticsClient
	Certificates                 *certificates.CertificatesClient
	ContainerApps                *containerapps.ContainerAppsClient
	ContainerAppsRevisions       *containerappsrevisions.ContainerAppsRevisionsClient
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	appServiceCertificateOrdersClient := appservicecertificateorders.NewAppServiceCertificateOrdersClientWithBaseURI(endpoint)
	configureAuthFunc(&appServiceCertificateOrdersClient.Client)

	appServiceEnvironmentsClient := appserviceenvironments.NewAppServiceEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&appServiceEnvironmentsClient.Client)

	appServicePlansClient := appserviceplans.NewAppServicePlansClientWithBaseURI(endpoint)
	configureAuthFunc(&appServicePlansClient.Client)

	certificateOrdersDiagnosticsClient := certificateordersdiagnostics.NewCertificateOrdersDiagnosticsClientWithBaseURI(endpoint)
	configureAuthFunc(&certificateOrdersDiagnosticsClient.Client)

	certificatesClient := certificates.NewCertificatesClientWithBaseURI(endpoint)
	configureAuthFunc(&certificatesClient.Client)

	containerAppsClient := containerapps.NewContainerAppsClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsClient.Client)

	containerAppsRevisionsClient := containerappsrevisions.NewContainerAppsRevisionsClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsRevisionsClient.Client)

	deletedWebAppsClient := deletedwebapps.NewDeletedWebAppsClientWithBaseURI(endpoint)
	configureAuthFunc(&deletedWebAppsClient.Client)

	diagnosticsClient := diagnostics.NewDiagnosticsClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticsClient.Client)

	domainsClient := domains.NewDomainsClientWithBaseURI(endpoint)
	configureAuthFunc(&domainsClient.Client)

	globalClient := global.NewGlobalClientWithBaseURI(endpoint)
	configureAuthFunc(&globalClient.Client)

	kubeEnvironmentsClient := kubeenvironments.NewKubeEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&kubeEnvironmentsClient.Client)

	providerClient := provider.NewProviderClientWithBaseURI(endpoint)
	configureAuthFunc(&providerClient.Client)

	recommendationsClient := recommendations.NewRecommendationsClientWithBaseURI(endpoint)
	configureAuthFunc(&recommendationsClient.Client)

	resourceHealthMetadataClient := resourcehealthmetadata.NewResourceHealthMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceHealthMetadataClient.Client)

	resourceProvidersClient := resourceproviders.NewResourceProvidersClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceProvidersClient.Client)

	staticSitesClient := staticsites.NewStaticSitesClientWithBaseURI(endpoint)
	configureAuthFunc(&staticSitesClient.Client)

	topLevelDomainsClient := topleveldomains.NewTopLevelDomainsClientWithBaseURI(endpoint)
	configureAuthFunc(&topLevelDomainsClient.Client)

	webAppsClient := webapps.NewWebAppsClientWithBaseURI(endpoint)
	configureAuthFunc(&webAppsClient.Client)

	workflowRunActionsClient := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowRunActionsClient.Client)

	workflowRunsClient := workflowruns.NewWorkflowRunsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowRunsClient.Client)

	workflowTriggerHistoriesClient := workflowtriggerhistories.NewWorkflowTriggerHistoriesClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowTriggerHistoriesClient.Client)

	workflowTriggersClient := workflowtriggers.NewWorkflowTriggersClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowTriggersClient.Client)

	workflowVersionsClient := workflowversions.NewWorkflowVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowVersionsClient.Client)

	workflowsClient := workflows.NewWorkflowsClientWithBaseURI(endpoint)
	configureAuthFunc(&workflowsClient.Client)

	return Client{
		AppServiceCertificateOrders:  &appServiceCertificateOrdersClient,
		AppServiceEnvironments:       &appServiceEnvironmentsClient,
		AppServicePlans:              &appServicePlansClient,
		CertificateOrdersDiagnostics: &certificateOrdersDiagnosticsClient,
		Certificates:                 &certificatesClient,
		ContainerApps:                &containerAppsClient,
		ContainerAppsRevisions:       &containerAppsRevisionsClient,
		DeletedWebApps:               &deletedWebAppsClient,
		Diagnostics:                  &diagnosticsClient,
		Domains:                      &domainsClient,
		Global:                       &globalClient,
		KubeEnvironments:             &kubeEnvironmentsClient,
		Provider:                     &providerClient,
		Recommendations:              &recommendationsClient,
		ResourceHealthMetadata:       &resourceHealthMetadataClient,
		ResourceProviders:            &resourceProvidersClient,
		StaticSites:                  &staticSitesClient,
		TopLevelDomains:              &topLevelDomainsClient,
		WebApps:                      &webAppsClient,
		WorkflowRunActions:           &workflowRunActionsClient,
		WorkflowRuns:                 &workflowRunsClient,
		WorkflowTriggerHistories:     &workflowTriggerHistoriesClient,
		WorkflowTriggers:             &workflowTriggersClient,
		WorkflowVersions:             &workflowVersionsClient,
		Workflows:                    &workflowsClient,
	}
}
