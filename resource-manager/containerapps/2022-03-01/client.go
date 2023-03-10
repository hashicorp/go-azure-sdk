package v2022_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/certificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/containerapps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/containerappsauthconfigs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/containerappsrevisionreplicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/containerappsrevisions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/containerappssourcecontrols"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/daprcomponents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/managedenvironments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/containerapps/2022-03-01/managedenvironmentsstorages"
)

type Client struct {
	Certificates                  *certificates.CertificatesClient
	ContainerApps                 *containerapps.ContainerAppsClient
	ContainerAppsAuthConfigs      *containerappsauthconfigs.ContainerAppsAuthConfigsClient
	ContainerAppsRevisionReplicas *containerappsrevisionreplicas.ContainerAppsRevisionReplicasClient
	ContainerAppsRevisions        *containerappsrevisions.ContainerAppsRevisionsClient
	ContainerAppsSourceControls   *containerappssourcecontrols.ContainerAppsSourceControlsClient
	DaprComponents                *daprcomponents.DaprComponentsClient
	ManagedEnvironments           *managedenvironments.ManagedEnvironmentsClient
	ManagedEnvironmentsStorages   *managedenvironmentsstorages.ManagedEnvironmentsStoragesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	certificatesClient := certificates.NewCertificatesClientWithBaseURI(endpoint)
	configureAuthFunc(&certificatesClient.Client)

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

	managedEnvironmentsClient := managedenvironments.NewManagedEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedEnvironmentsClient.Client)

	managedEnvironmentsStoragesClient := managedenvironmentsstorages.NewManagedEnvironmentsStoragesClientWithBaseURI(endpoint)
	configureAuthFunc(&managedEnvironmentsStoragesClient.Client)

	return Client{
		Certificates:                  &certificatesClient,
		ContainerApps:                 &containerAppsClient,
		ContainerAppsAuthConfigs:      &containerAppsAuthConfigsClient,
		ContainerAppsRevisionReplicas: &containerAppsRevisionReplicasClient,
		ContainerAppsRevisions:        &containerAppsRevisionsClient,
		ContainerAppsSourceControls:   &containerAppsSourceControlsClient,
		DaprComponents:                &daprComponentsClient,
		ManagedEnvironments:           &managedEnvironmentsClient,
		ManagedEnvironmentsStorages:   &managedEnvironmentsStoragesClient,
	}
}
