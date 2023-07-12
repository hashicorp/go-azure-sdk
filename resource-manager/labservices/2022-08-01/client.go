package v2022_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/image"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/lab"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/labplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/user"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2022-08-01/virtualmachine"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	imageClient, err := image.NewImageClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Image client: %+v", err)
	}
	configureFunc(imageClient.Client)

	labClient, err := lab.NewLabClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Lab client: %+v", err)
	}
	configureFunc(labClient.Client)

	labPlanClient, err := labplan.NewLabPlanClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LabPlan client: %+v", err)
	}
	configureFunc(labPlanClient.Client)

	scheduleClient, err := schedule.NewScheduleClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %+v", err)
	}
	configureFunc(scheduleClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	userClient, err := user.NewUserClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building User client: %+v", err)
	}
	configureFunc(userClient.Client)

	virtualMachineClient, err := virtualmachine.NewVirtualMachineClientWithBaseURI(api)
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
