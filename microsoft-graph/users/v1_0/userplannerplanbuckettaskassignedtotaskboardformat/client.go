package userplannerplanbuckettaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanBucketTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanBucketTaskAssignedToTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanBucketTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplanbuckettaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanBucketTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerPlanBucketTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
