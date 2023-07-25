package v2018_09_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	armTemplatesClient, err := armtemplates.NewArmTemplatesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ArmTemplates client: %+v", err)
	}
	configureFunc(armTemplatesClient.Client)

	artifactSourcesClient, err := artifactsources.NewArtifactSourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ArtifactSources client: %+v", err)
	}
	configureFunc(artifactSourcesClient.Client)

	artifactsClient, err := artifacts.NewArtifactsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Artifacts client: %+v", err)
	}
	configureFunc(artifactsClient.Client)

	costsClient, err := costs.NewCostsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Costs client: %+v", err)
	}
	configureFunc(costsClient.Client)

	customImagesClient, err := customimages.NewCustomImagesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CustomImages client: %+v", err)
	}
	configureFunc(customImagesClient.Client)

	disksClient, err := disks.NewDisksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Disks client: %+v", err)
	}
	configureFunc(disksClient.Client)

	environmentsClient, err := environments.NewEnvironmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Environments client: %+v", err)
	}
	configureFunc(environmentsClient.Client)

	formulasClient, err := formulas.NewFormulasClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Formulas client: %+v", err)
	}
	configureFunc(formulasClient.Client)

	galleryImagesClient, err := galleryimages.NewGalleryImagesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImages client: %+v", err)
	}
	configureFunc(galleryImagesClient.Client)

	globalSchedulesClient, err := globalschedules.NewGlobalSchedulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building GlobalSchedules client: %+v", err)
	}
	configureFunc(globalSchedulesClient.Client)

	labsClient, err := labs.NewLabsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Labs client: %+v", err)
	}
	configureFunc(labsClient.Client)

	notificationChannelsClient, err := notificationchannels.NewNotificationChannelsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building NotificationChannels client: %+v", err)
	}
	configureFunc(notificationChannelsClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	policiesClient, err := policies.NewPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Policies client: %+v", err)
	}
	configureFunc(policiesClient.Client)

	policySetsClient, err := policysets.NewPolicySetsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PolicySets client: %+v", err)
	}
	configureFunc(policySetsClient.Client)

	schedulesClient, err := schedules.NewSchedulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Schedules client: %+v", err)
	}
	configureFunc(schedulesClient.Client)

	secretsClient, err := secrets.NewSecretsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Secrets client: %+v", err)
	}
	configureFunc(secretsClient.Client)

	serviceFabricSchedulesClient, err := servicefabricschedules.NewServiceFabricSchedulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServiceFabricSchedules client: %+v", err)
	}
	configureFunc(serviceFabricSchedulesClient.Client)

	serviceFabricsClient, err := servicefabrics.NewServiceFabricsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServiceFabrics client: %+v", err)
	}
	configureFunc(serviceFabricsClient.Client)

	serviceRunnersClient, err := servicerunners.NewServiceRunnersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServiceRunners client: %+v", err)
	}
	configureFunc(serviceRunnersClient.Client)

	usersClient, err := users.NewUsersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Users client: %+v", err)
	}
	configureFunc(usersClient.Client)

	virtualMachineSchedulesClient, err := virtualmachineschedules.NewVirtualMachineSchedulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineSchedules client: %+v", err)
	}
	configureFunc(virtualMachineSchedulesClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	virtualNetworksClient, err := virtualnetworks.NewVirtualNetworksClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworks client: %+v", err)
	}
	configureFunc(virtualNetworksClient.Client)

	return &Client{
		ArmTemplates:            armTemplatesClient,
		ArtifactSources:         artifactSourcesClient,
		Artifacts:               artifactsClient,
		Costs:                   costsClient,
		CustomImages:            customImagesClient,
		Disks:                   disksClient,
		Environments:            environmentsClient,
		Formulas:                formulasClient,
		GalleryImages:           galleryImagesClient,
		GlobalSchedules:         globalSchedulesClient,
		Labs:                    labsClient,
		NotificationChannels:    notificationChannelsClient,
		Operations:              operationsClient,
		Policies:                policiesClient,
		PolicySets:              policySetsClient,
		Schedules:               schedulesClient,
		Secrets:                 secretsClient,
		ServiceFabricSchedules:  serviceFabricSchedulesClient,
		ServiceFabrics:          serviceFabricsClient,
		ServiceRunners:          serviceRunnersClient,
		Users:                   usersClient,
		VirtualMachineSchedules: virtualMachineSchedulesClient,
		VirtualMachines:         virtualMachinesClient,
		VirtualNetworks:         virtualNetworksClient,
	}, nil
}
