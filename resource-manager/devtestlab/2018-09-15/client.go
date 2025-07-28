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
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/dtls"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/environments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/formulas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/galleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/globalschedules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/labs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/devtestlab/2018-09-15/notificationchannels"
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
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ArmTemplates            *armtemplates.ArmTemplatesClient
	ArtifactSources         *artifactsources.ArtifactSourcesClient
	Artifacts               *artifacts.ArtifactsClient
	Costs                   *costs.CostsClient
	CustomImages            *customimages.CustomImagesClient
	DTLS                    *dtls.DTLSClient
	Disks                   *disks.DisksClient
	Environments            *environments.EnvironmentsClient
	Formulas                *formulas.FormulasClient
	GalleryImages           *galleryimages.GalleryImagesClient
	GlobalSchedules         *globalschedules.GlobalSchedulesClient
	Labs                    *labs.LabsClient
	NotificationChannels    *notificationchannels.NotificationChannelsClient
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	armTemplatesClient, err := armtemplates.NewArmTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ArmTemplates client: %+v", err)
	}
	configureFunc(armTemplatesClient.Client)

	artifactSourcesClient, err := artifactsources.NewArtifactSourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ArtifactSources client: %+v", err)
	}
	configureFunc(artifactSourcesClient.Client)

	artifactsClient, err := artifacts.NewArtifactsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Artifacts client: %+v", err)
	}
	configureFunc(artifactsClient.Client)

	costsClient, err := costs.NewCostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Costs client: %+v", err)
	}
	configureFunc(costsClient.Client)

	customImagesClient, err := customimages.NewCustomImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomImages client: %+v", err)
	}
	configureFunc(customImagesClient.Client)

	dTLSClient, err := dtls.NewDTLSClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DTLS client: %+v", err)
	}
	configureFunc(dTLSClient.Client)

	disksClient, err := disks.NewDisksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Disks client: %+v", err)
	}
	configureFunc(disksClient.Client)

	environmentsClient, err := environments.NewEnvironmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Environments client: %+v", err)
	}
	configureFunc(environmentsClient.Client)

	formulasClient, err := formulas.NewFormulasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Formulas client: %+v", err)
	}
	configureFunc(formulasClient.Client)

	galleryImagesClient, err := galleryimages.NewGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImages client: %+v", err)
	}
	configureFunc(galleryImagesClient.Client)

	globalSchedulesClient, err := globalschedules.NewGlobalSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GlobalSchedules client: %+v", err)
	}
	configureFunc(globalSchedulesClient.Client)

	labsClient, err := labs.NewLabsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Labs client: %+v", err)
	}
	configureFunc(labsClient.Client)

	notificationChannelsClient, err := notificationchannels.NewNotificationChannelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotificationChannels client: %+v", err)
	}
	configureFunc(notificationChannelsClient.Client)

	policiesClient, err := policies.NewPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Policies client: %+v", err)
	}
	configureFunc(policiesClient.Client)

	policySetsClient, err := policysets.NewPolicySetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicySets client: %+v", err)
	}
	configureFunc(policySetsClient.Client)

	schedulesClient, err := schedules.NewSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Schedules client: %+v", err)
	}
	configureFunc(schedulesClient.Client)

	secretsClient, err := secrets.NewSecretsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Secrets client: %+v", err)
	}
	configureFunc(secretsClient.Client)

	serviceFabricSchedulesClient, err := servicefabricschedules.NewServiceFabricSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceFabricSchedules client: %+v", err)
	}
	configureFunc(serviceFabricSchedulesClient.Client)

	serviceFabricsClient, err := servicefabrics.NewServiceFabricsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceFabrics client: %+v", err)
	}
	configureFunc(serviceFabricsClient.Client)

	serviceRunnersClient, err := servicerunners.NewServiceRunnersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceRunners client: %+v", err)
	}
	configureFunc(serviceRunnersClient.Client)

	usersClient, err := users.NewUsersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Users client: %+v", err)
	}
	configureFunc(usersClient.Client)

	virtualMachineSchedulesClient, err := virtualmachineschedules.NewVirtualMachineSchedulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineSchedules client: %+v", err)
	}
	configureFunc(virtualMachineSchedulesClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	virtualNetworksClient, err := virtualnetworks.NewVirtualNetworksClientWithBaseURI(sdkApi)
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
		DTLS:                    dTLSClient,
		Disks:                   disksClient,
		Environments:            environmentsClient,
		Formulas:                formulasClient,
		GalleryImages:           galleryImagesClient,
		GlobalSchedules:         globalSchedulesClient,
		Labs:                    labsClient,
		NotificationChannels:    notificationChannelsClient,
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
