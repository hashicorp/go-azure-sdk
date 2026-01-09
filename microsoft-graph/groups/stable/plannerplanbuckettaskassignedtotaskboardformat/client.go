package plannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbuckettaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanBucketTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
