package v2022_11_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/availableworkloadprofiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/billingmeters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/connectedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/connectedenvironmentsstorages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerappsauthconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerappsrevisionreplicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerappsrevisions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/containerappssourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/daprcomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/diagnostics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedcertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-11-01-preview/managedenvironmentsstorages"
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
	ContainerAppsSourceControls   *containerappssourcecontrols.ContainerAppsSourceControlsClient
	DaprComponents                *daprcomponents.DaprComponentsClient
	Diagnostics                   *diagnostics.DiagnosticsClient
	Jobs                          *jobs.JobsClient
	ManagedCertificates           *managedcertificates.ManagedCertificatesClient
	ManagedEnvironments           *managedenvironments.ManagedEnvironmentsClient
	ManagedEnvironmentsStorages   *managedenvironmentsstorages.ManagedEnvironmentsStoragesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	availableWorkloadProfilesClient := availableworkloadprofiles.NewAvailableWorkloadProfilesClientWithBaseURI(endpoint)
	configureAuthFunc(&availableWorkloadProfilesClient.Client)

	billingMetersClient := billingmeters.NewBillingMetersClientWithBaseURI(endpoint)
	configureAuthFunc(&billingMetersClient.Client)

	certificatesClient := certificates.NewCertificatesClientWithBaseURI(endpoint)
	configureAuthFunc(&certificatesClient.Client)

	connectedEnvironmentsClient := connectedenvironments.NewConnectedEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&connectedEnvironmentsClient.Client)

	connectedEnvironmentsStoragesClient := connectedenvironmentsstorages.NewConnectedEnvironmentsStoragesClientWithBaseURI(endpoint)
	configureAuthFunc(&connectedEnvironmentsStoragesClient.Client)

	containerAppsAuthConfigsClient := containerappsauthconfigs.NewContainerAppsAuthConfigsClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsAuthConfigsClient.Client)

	containerAppsClient := containerapps.NewContainerAppsClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsClient.Client)

	containerAppsRevisionReplicasClient := containerappsrevisionreplicas.NewContainerAppsRevisionReplicasClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsRevisionReplicasClient.Client)

	containerAppsRevisionsClient := containerappsrevisions.NewContainerAppsRevisionsClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsRevisionsClient.Client)

	containerAppsSourceControlsClient := containerappssourcecontrols.NewContainerAppsSourceControlsClientWithBaseURI(endpoint)
	configureAuthFunc(&containerAppsSourceControlsClient.Client)

	daprComponentsClient := daprcomponents.NewDaprComponentsClientWithBaseURI(endpoint)
	configureAuthFunc(&daprComponentsClient.Client)

	diagnosticsClient := diagnostics.NewDiagnosticsClientWithBaseURI(endpoint)
	configureAuthFunc(&diagnosticsClient.Client)

	jobsClient := jobs.NewJobsClientWithBaseURI(endpoint)
	configureAuthFunc(&jobsClient.Client)

	managedCertificatesClient := managedcertificates.NewManagedCertificatesClientWithBaseURI(endpoint)
	configureAuthFunc(&managedCertificatesClient.Client)

	managedEnvironmentsClient := managedenvironments.NewManagedEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedEnvironmentsClient.Client)

	managedEnvironmentsStoragesClient := managedenvironmentsstorages.NewManagedEnvironmentsStoragesClientWithBaseURI(endpoint)
	configureAuthFunc(&managedEnvironmentsStoragesClient.Client)

	return Client{
		AvailableWorkloadProfiles:     &availableWorkloadProfilesClient,
		BillingMeters:                 &billingMetersClient,
		Certificates:                  &certificatesClient,
		ConnectedEnvironments:         &connectedEnvironmentsClient,
		ConnectedEnvironmentsStorages: &connectedEnvironmentsStoragesClient,
		ContainerApps:                 &containerAppsClient,
		ContainerAppsAuthConfigs:      &containerAppsAuthConfigsClient,
		ContainerAppsRevisionReplicas: &containerAppsRevisionReplicasClient,
		ContainerAppsRevisions:        &containerAppsRevisionsClient,
		ContainerAppsSourceControls:   &containerAppsSourceControlsClient,
		DaprComponents:                &daprComponentsClient,
		Diagnostics:                   &diagnosticsClient,
		Jobs:                          &jobsClient,
		ManagedCertificates:           &managedCertificatesClient,
		ManagedEnvironments:           &managedEnvironmentsClient,
		ManagedEnvironmentsStorages:   &managedEnvironmentsStoragesClient,
	}
}
