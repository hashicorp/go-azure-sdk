package v2024_02_02_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/appresiliency"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/availableworkloadprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/billingmeters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/builders"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/builds"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/connectedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/connectedenvironmentsstorages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsauthconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsbuilds"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappspatches"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsrevisionreplicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappsrevisions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappssessionpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/containerappssourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/daprcomponentresiliencypolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/daprcomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/daprsubscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/diagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/dotnetcomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/functionsextension"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/javacomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/logicapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/managedcertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/managedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/managedenvironmentsstorages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2024-02-02-preview/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AppResiliency                   *appresiliency.AppResiliencyClient
	AvailableWorkloadProfiles       *availableworkloadprofiles.AvailableWorkloadProfilesClient
	BillingMeters                   *billingmeters.BillingMetersClient
	Builders                        *builders.BuildersClient
	Builds                          *builds.BuildsClient
	Certificates                    *certificates.CertificatesClient
	ConnectedEnvironments           *connectedenvironments.ConnectedEnvironmentsClient
	ConnectedEnvironmentsStorages   *connectedenvironmentsstorages.ConnectedEnvironmentsStoragesClient
	ContainerApps                   *containerapps.ContainerAppsClient
	ContainerAppsAuthConfigs        *containerappsauthconfigs.ContainerAppsAuthConfigsClient
	ContainerAppsBuilds             *containerappsbuilds.ContainerAppsBuildsClient
	ContainerAppsPatches            *containerappspatches.ContainerAppsPatchesClient
	ContainerAppsRevisionReplicas   *containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient
	ContainerAppsRevisions          *containerappsrevisions.ContainerAppsRevisionsClient
	ContainerAppsSessionPools       *containerappssessionpools.ContainerAppsSessionPoolsClient
	ContainerAppsSourceControls     *containerappssourcecontrols.ContainerAppsSourceControlsClient
	DaprComponentResiliencyPolicies *daprcomponentresiliencypolicies.DaprComponentResiliencyPoliciesClient
	DaprComponents                  *daprcomponents.DaprComponentsClient
	DaprSubscriptions               *daprsubscriptions.DaprSubscriptionsClient
	Diagnostics                     *diagnostics.DiagnosticsClient
	DotNetComponents                *dotnetcomponents.DotNetComponentsClient
	FunctionsExtension              *functionsextension.FunctionsExtensionClient
	JavaComponents                  *javacomponents.JavaComponentsClient
	Jobs                            *jobs.JobsClient
	LogicApps                       *logicapps.LogicAppsClient
	ManagedCertificates             *managedcertificates.ManagedCertificatesClient
	ManagedEnvironments             *managedenvironments.ManagedEnvironmentsClient
	ManagedEnvironmentsStorages     *managedenvironmentsstorages.ManagedEnvironmentsStoragesClient
	PrivateEndpointConnections      *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources            *privatelinkresources.PrivateLinkResourcesClient
	Subscriptions                   *subscriptions.SubscriptionsClient
	Usages                          *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	appResiliencyClient, err := appresiliency.NewAppResiliencyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AppResiliency client: %+v", err)
	}
	configureFunc(appResiliencyClient.Client)

	availableWorkloadProfilesClient, err := availableworkloadprofiles.NewAvailableWorkloadProfilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailableWorkloadProfiles client: %+v", err)
	}
	configureFunc(availableWorkloadProfilesClient.Client)

	billingMetersClient, err := billingmeters.NewBillingMetersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BillingMeters client: %+v", err)
	}
	configureFunc(billingMetersClient.Client)

	buildersClient, err := builders.NewBuildersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Builders client: %+v", err)
	}
	configureFunc(buildersClient.Client)

	buildsClient, err := builds.NewBuildsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Builds client: %+v", err)
	}
	configureFunc(buildsClient.Client)

	certificatesClient, err := certificates.NewCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificates client: %+v", err)
	}
	configureFunc(certificatesClient.Client)

	connectedEnvironmentsClient, err := connectedenvironments.NewConnectedEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedEnvironments client: %+v", err)
	}
	configureFunc(connectedEnvironmentsClient.Client)

	connectedEnvironmentsStoragesClient, err := connectedenvironmentsstorages.NewConnectedEnvironmentsStoragesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectedEnvironmentsStorages client: %+v", err)
	}
	configureFunc(connectedEnvironmentsStoragesClient.Client)

	containerAppsAuthConfigsClient, err := containerappsauthconfigs.NewContainerAppsAuthConfigsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsAuthConfigs client: %+v", err)
	}
	configureFunc(containerAppsAuthConfigsClient.Client)

	containerAppsBuildsClient, err := containerappsbuilds.NewContainerAppsBuildsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsBuilds client: %+v", err)
	}
	configureFunc(containerAppsBuildsClient.Client)

	containerAppsClient, err := containerapps.NewContainerAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerApps client: %+v", err)
	}
	configureFunc(containerAppsClient.Client)

	containerAppsPatchesClient, err := containerappspatches.NewContainerAppsPatchesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsPatches client: %+v", err)
	}
	configureFunc(containerAppsPatchesClient.Client)

	containerAppsRevisionReplicasClient, err := containerappsrevisionreplicas.NewContainerAppsRevisionReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsRevisionReplicas client: %+v", err)
	}
	configureFunc(containerAppsRevisionReplicasClient.Client)

	containerAppsRevisionsClient, err := containerappsrevisions.NewContainerAppsRevisionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsRevisions client: %+v", err)
	}
	configureFunc(containerAppsRevisionsClient.Client)

	containerAppsSessionPoolsClient, err := containerappssessionpools.NewContainerAppsSessionPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsSessionPools client: %+v", err)
	}
	configureFunc(containerAppsSessionPoolsClient.Client)

	containerAppsSourceControlsClient, err := containerappssourcecontrols.NewContainerAppsSourceControlsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerAppsSourceControls client: %+v", err)
	}
	configureFunc(containerAppsSourceControlsClient.Client)

	daprComponentResiliencyPoliciesClient, err := daprcomponentresiliencypolicies.NewDaprComponentResiliencyPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DaprComponentResiliencyPolicies client: %+v", err)
	}
	configureFunc(daprComponentResiliencyPoliciesClient.Client)

	daprComponentsClient, err := daprcomponents.NewDaprComponentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DaprComponents client: %+v", err)
	}
	configureFunc(daprComponentsClient.Client)

	daprSubscriptionsClient, err := daprsubscriptions.NewDaprSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DaprSubscriptions client: %+v", err)
	}
	configureFunc(daprSubscriptionsClient.Client)

	diagnosticsClient, err := diagnostics.NewDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Diagnostics client: %+v", err)
	}
	configureFunc(diagnosticsClient.Client)

	dotNetComponentsClient, err := dotnetcomponents.NewDotNetComponentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DotNetComponents client: %+v", err)
	}
	configureFunc(dotNetComponentsClient.Client)

	functionsExtensionClient, err := functionsextension.NewFunctionsExtensionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FunctionsExtension client: %+v", err)
	}
	configureFunc(functionsExtensionClient.Client)

	javaComponentsClient, err := javacomponents.NewJavaComponentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JavaComponents client: %+v", err)
	}
	configureFunc(javaComponentsClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	logicAppsClient, err := logicapps.NewLogicAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogicApps client: %+v", err)
	}
	configureFunc(logicAppsClient.Client)

	managedCertificatesClient, err := managedcertificates.NewManagedCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedCertificates client: %+v", err)
	}
	configureFunc(managedCertificatesClient.Client)

	managedEnvironmentsClient, err := managedenvironments.NewManagedEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedEnvironments client: %+v", err)
	}
	configureFunc(managedEnvironmentsClient.Client)

	managedEnvironmentsStoragesClient, err := managedenvironmentsstorages.NewManagedEnvironmentsStoragesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedEnvironmentsStorages client: %+v", err)
	}
	configureFunc(managedEnvironmentsStoragesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	subscriptionsClient, err := subscriptions.NewSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscriptions client: %+v", err)
	}
	configureFunc(subscriptionsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	return &Client{
		AppResiliency:                   appResiliencyClient,
		AvailableWorkloadProfiles:       availableWorkloadProfilesClient,
		BillingMeters:                   billingMetersClient,
		Builders:                        buildersClient,
		Builds:                          buildsClient,
		Certificates:                    certificatesClient,
		ConnectedEnvironments:           connectedEnvironmentsClient,
		ConnectedEnvironmentsStorages:   connectedEnvironmentsStoragesClient,
		ContainerApps:                   containerAppsClient,
		ContainerAppsAuthConfigs:        containerAppsAuthConfigsClient,
		ContainerAppsBuilds:             containerAppsBuildsClient,
		ContainerAppsPatches:            containerAppsPatchesClient,
		ContainerAppsRevisionReplicas:   containerAppsRevisionReplicasClient,
		ContainerAppsRevisions:          containerAppsRevisionsClient,
		ContainerAppsSessionPools:       containerAppsSessionPoolsClient,
		ContainerAppsSourceControls:     containerAppsSourceControlsClient,
		DaprComponentResiliencyPolicies: daprComponentResiliencyPoliciesClient,
		DaprComponents:                  daprComponentsClient,
		DaprSubscriptions:               daprSubscriptionsClient,
		Diagnostics:                     diagnosticsClient,
		DotNetComponents:                dotNetComponentsClient,
		FunctionsExtension:              functionsExtensionClient,
		JavaComponents:                  javaComponentsClient,
		Jobs:                            jobsClient,
		LogicApps:                       logicAppsClient,
		ManagedCertificates:             managedCertificatesClient,
		ManagedEnvironments:             managedEnvironmentsClient,
		ManagedEnvironmentsStorages:     managedEnvironmentsStoragesClient,
		PrivateEndpointConnections:      privateEndpointConnectionsClient,
		PrivateLinkResources:            privateLinkResourcesClient,
		Subscriptions:                   subscriptionsClient,
		Usages:                          usagesClient,
	}, nil
}
