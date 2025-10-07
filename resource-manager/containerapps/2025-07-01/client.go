package v2025_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/availableworkloadprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/billingmeters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/connectedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/connectedenvironmentsstorages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/containerapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/containerappsauthconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/containerappsrevisionreplicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/containerappsrevisions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/containerappssessionpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/containerappssourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/daprcomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/diagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/httprouteconfig"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/javacomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/logicapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/maintenanceconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/managedcertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/managedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/managedenvironmentsstorages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/subscriptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2025-07-01/usages"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailableWorkloadProfiles     *availableworkloadprofiles.AvailableWorkloadProfilesClient
	BillingMeters                 *billingmeters.BillingMetersClient
	Certificates                  *certificates.CertificatesClient
	ConnectedEnvironments         *connectedenvironments.ConnectedEnvironmentsClient
	ConnectedEnvironmentsStorages *connectedenvironmentsstorages.ConnectedEnvironmentsStoragesClient
	ContainerApps                 *containerapps.ContainerAppsClient
	ContainerAppsAuthConfigs      *containerappsauthconfigs.ContainerAppsAuthConfigsClient
	ContainerAppsRevisionReplicas *containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient
	ContainerAppsRevisions        *containerappsrevisions.ContainerAppsRevisionsClient
	ContainerAppsSessionPools     *containerappssessionpools.ContainerAppsSessionPoolsClient
	ContainerAppsSourceControls   *containerappssourcecontrols.ContainerAppsSourceControlsClient
	DaprComponents                *daprcomponents.DaprComponentsClient
	Diagnostics                   *diagnostics.DiagnosticsClient
	HTTPRouteConfig               *httprouteconfig.HTTPRouteConfigClient
	JavaComponents                *javacomponents.JavaComponentsClient
	Jobs                          *jobs.JobsClient
	LogicApps                     *logicapps.LogicAppsClient
	MaintenanceConfigurations     *maintenanceconfigurations.MaintenanceConfigurationsClient
	ManagedCertificates           *managedcertificates.ManagedCertificatesClient
	ManagedEnvironments           *managedenvironments.ManagedEnvironmentsClient
	ManagedEnvironmentsStorages   *managedenvironmentsstorages.ManagedEnvironmentsStoragesClient
	PrivateEndpointConnections    *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources          *privatelinkresources.PrivateLinkResourcesClient
	Subscriptions                 *subscriptions.SubscriptionsClient
	Usages                        *usages.UsagesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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

	containerAppsClient, err := containerapps.NewContainerAppsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContainerApps client: %+v", err)
	}
	configureFunc(containerAppsClient.Client)

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

	daprComponentsClient, err := daprcomponents.NewDaprComponentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DaprComponents client: %+v", err)
	}
	configureFunc(daprComponentsClient.Client)

	diagnosticsClient, err := diagnostics.NewDiagnosticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Diagnostics client: %+v", err)
	}
	configureFunc(diagnosticsClient.Client)

	hTTPRouteConfigClient, err := httprouteconfig.NewHTTPRouteConfigClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HTTPRouteConfig client: %+v", err)
	}
	configureFunc(hTTPRouteConfigClient.Client)

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

	maintenanceConfigurationsClient, err := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceConfigurations client: %+v", err)
	}
	configureFunc(maintenanceConfigurationsClient.Client)

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
		AvailableWorkloadProfiles:     availableWorkloadProfilesClient,
		BillingMeters:                 billingMetersClient,
		Certificates:                  certificatesClient,
		ConnectedEnvironments:         connectedEnvironmentsClient,
		ConnectedEnvironmentsStorages: connectedEnvironmentsStoragesClient,
		ContainerApps:                 containerAppsClient,
		ContainerAppsAuthConfigs:      containerAppsAuthConfigsClient,
		ContainerAppsRevisionReplicas: containerAppsRevisionReplicasClient,
		ContainerAppsRevisions:        containerAppsRevisionsClient,
		ContainerAppsSessionPools:     containerAppsSessionPoolsClient,
		ContainerAppsSourceControls:   containerAppsSourceControlsClient,
		DaprComponents:                daprComponentsClient,
		Diagnostics:                   diagnosticsClient,
		HTTPRouteConfig:               hTTPRouteConfigClient,
		JavaComponents:                javaComponentsClient,
		Jobs:                          jobsClient,
		LogicApps:                     logicAppsClient,
		MaintenanceConfigurations:     maintenanceConfigurationsClient,
		ManagedCertificates:           managedCertificatesClient,
		ManagedEnvironments:           managedEnvironmentsClient,
		ManagedEnvironmentsStorages:   managedEnvironmentsStoragesClient,
		PrivateEndpointConnections:    privateEndpointConnectionsClient,
		PrivateLinkResources:          privateLinkResourcesClient,
		Subscriptions:                 subscriptionsClient,
		Usages:                        usagesClient,
	}, nil
}
