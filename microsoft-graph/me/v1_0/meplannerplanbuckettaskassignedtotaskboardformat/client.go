package meplannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanBucketTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanBucketTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplanbuckettaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanBucketTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &MePlannerPlanBucketTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
