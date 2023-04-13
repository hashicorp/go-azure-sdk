package v2018_09_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/armtemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/artifacts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/artifactsources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/costs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/customimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/disks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/environments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/formulas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/galleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/globalschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/labs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/notificationchannels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/policies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/policysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/schedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/secrets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicefabrics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicefabricschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/servicerunners"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/users"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/virtualmachineschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/virtualnetworks"
)

type Client struct {
	ArmTemplates            *armtemplates.ArmTemplatesClient
	ArtifactSources         *artifactsources.ArtifactSourcesClient
	Artifacts               *artifacts.ArtifactsClient
	Costs                   *costs.CostsClient
	CustomImages            *customimages.CustomImagesClient
	Disks                   *disks.DisksClient
	Environments            *environments.EnvironmentsClient
	Formulas                *formulas.FormulasClient
	GalleryImages           *galleryimages.GalleryImagesClient
	GlobalSchedules         *globalschedules.GlobalSchedulesClient
	Labs                    *labs.LabsClient
	NotificationChannels    *notificationchannels.NotificationChannelsClient
	Operations              *operations.OperationsClient
	Policies                *policies.PoliciesClient
	PolicySets              *policysets.PolicySetsClient
	Schedules               *schedules.SchedulesClient
	Secrets                 *secrets.SecretsClient
	ServiceFabricSchedules  *servicefabricschedules.ServiceFabricSchedulesClient
	ServiceFabrics          *servicefabrics.ServiceFabricsClient
	ServiceRunners          *servicerunners.ServiceRunnersClient
	Users                   *users.UsersClient
	VirtualMachineSchedules *virtualmachineschedules.VirtualMachineSchedulesClient
	VirtualMachines         *virtualmachines.VirtualMachinesClient
	VirtualNetworks         *virtualnetworks.VirtualNetworksClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	armTemplatesClient := armtemplates.NewArmTemplatesClientWithBaseURI(endpoint)
	configureAuthFunc(&armTemplatesClient.Client)

	artifactSourcesClient := artifactsources.NewArtifactSourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&artifactSourcesClient.Client)

	artifactsClient := artifacts.NewArtifactsClientWithBaseURI(endpoint)
	configureAuthFunc(&artifactsClient.Client)

	costsClient := costs.NewCostsClientWithBaseURI(endpoint)
	configureAuthFunc(&costsClient.Client)

	customImagesClient := customimages.NewCustomImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&customImagesClient.Client)

	disksClient := disks.NewDisksClientWithBaseURI(endpoint)
	configureAuthFunc(&disksClient.Client)

	environmentsClient := environments.NewEnvironmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&environmentsClient.Client)

	formulasClient := formulas.NewFormulasClientWithBaseURI(endpoint)
	configureAuthFunc(&formulasClient.Client)

	galleryImagesClient := galleryimages.NewGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryImagesClient.Client)

	globalSchedulesClient := globalschedules.NewGlobalSchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&globalSchedulesClient.Client)

	labsClient := labs.NewLabsClientWithBaseURI(endpoint)
	configureAuthFunc(&labsClient.Client)

	notificationChannelsClient := notificationchannels.NewNotificationChannelsClientWithBaseURI(endpoint)
	configureAuthFunc(&notificationChannelsClient.Client)

	operationsClient := operations.NewOperationsClientWithBaseURI(endpoint)
	configureAuthFunc(&operationsClient.Client)

	policiesClient := policies.NewPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&policiesClient.Client)

	policySetsClient := policysets.NewPolicySetsClientWithBaseURI(endpoint)
	configureAuthFunc(&policySetsClient.Client)

	schedulesClient := schedules.NewSchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&schedulesClient.Client)

	secretsClient := secrets.NewSecretsClientWithBaseURI(endpoint)
	configureAuthFunc(&secretsClient.Client)

	serviceFabricSchedulesClient := servicefabricschedules.NewServiceFabricSchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceFabricSchedulesClient.Client)

	serviceFabricsClient := servicefabrics.NewServiceFabricsClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceFabricsClient.Client)

	serviceRunnersClient := servicerunners.NewServiceRunnersClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceRunnersClient.Client)

	usersClient := users.NewUsersClientWithBaseURI(endpoint)
	configureAuthFunc(&usersClient.Client)

	virtualMachineSchedulesClient := virtualmachineschedules.NewVirtualMachineSchedulesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineSchedulesClient.Client)

	virtualMachinesClient := virtualmachines.NewVirtualMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachinesClient.Client)

	virtualNetworksClient := virtualnetworks.NewVirtualNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworksClient.Client)

	return Client{
		ArmTemplates:            &armTemplatesClient,
		ArtifactSources:         &artifactSourcesClient,
		Artifacts:               &artifactsClient,
		Costs:                   &costsClient,
		CustomImages:            &customImagesClient,
		Disks:                   &disksClient,
		Environments:            &environmentsClient,
		Formulas:                &formulasClient,
		GalleryImages:           &galleryImagesClient,
		GlobalSchedules:         &globalSchedulesClient,
		Labs:                    &labsClient,
		NotificationChannels:    &notificationChannelsClient,
		Operations:              &operationsClient,
		Policies:                &policiesClient,
		PolicySets:              &policySetsClient,
		Schedules:               &schedulesClient,
		Secrets:                 &secretsClient,
		ServiceFabricSchedules:  &serviceFabricSchedulesClient,
		ServiceFabrics:          &serviceFabricsClient,
		ServiceRunners:          &serviceRunnersClient,
		Users:                   &usersClient,
		VirtualMachineSchedules: &virtualMachineSchedulesClient,
		VirtualMachines:         &virtualMachinesClient,
		VirtualNetworks:         &virtualNetworksClient,
	}
}
