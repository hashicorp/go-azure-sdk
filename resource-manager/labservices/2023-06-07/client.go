package v2023_06_07

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/image"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/lab"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/labplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/user"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2023-06-07/virtualmachine"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Image          *image.ImageClient
	Lab            *lab.LabClient
	LabPlan        *labplan.LabPlanClient
	Schedule       *schedule.ScheduleClient
	Skus           *skus.SkusClient
	Usages         *usages.UsagesClient
	User           *user.UserClient
	VirtualMachine *virtualmachine.VirtualMachineClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	imageClient, err := image.NewImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Image client: %+v", err)
	}
	configureFunc(imageClient.Client)

	labClient, err := lab.NewLabClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Lab client: %+v", err)
	}
	configureFunc(labClient.Client)

	labPlanClient, err := labplan.NewLabPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LabPlan client: %+v", err)
	}
	configureFunc(labPlanClient.Client)

	scheduleClient, err := schedule.NewScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %+v", err)
	}
	configureFunc(scheduleClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building User client: %+v", err)
	}
	configureFunc(userClient.Client)

	virtualMachineClient, err := virtualmachine.NewVirtualMachineClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachine client: %+v", err)
	}
	configureFunc(virtualMachineClient.Client)

	return &Client{
		Image:          imageClient,
		Lab:            labClient,
		LabPlan:        labPlanClient,
		Schedule:       scheduleClient,
		Skus:           skusClient,
		Usages:         usagesClient,
		User:           userClient,
		VirtualMachine: virtualMachineClient,
	}, nil
}
