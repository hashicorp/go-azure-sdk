package groupplannerplantaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplantaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &GroupPlannerPlanTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
