package plannerplantaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplantaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
