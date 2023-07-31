package v2021_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/image"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/lab"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/labplan"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/user"
	"github.com/hashicorp/go-azure-sdk/resource-manager/labservices/2021-10-01-preview/virtualmachine"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Image          *image.ImageClient
	Lab            *lab.LabClient
	LabPlan        *labplan.LabPlanClient
	Schedule       *schedule.ScheduleClient
	User           *user.UserClient
	VirtualMachine *virtualmachine.VirtualMachineClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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
		User:           userClient,
		VirtualMachine: virtualMachineClient,
	}, nil
}
