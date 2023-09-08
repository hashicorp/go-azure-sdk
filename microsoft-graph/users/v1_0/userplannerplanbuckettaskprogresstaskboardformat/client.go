package userplannerplanbuckettaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanBucketTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanBucketTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplanbuckettaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanBucketTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerPlanBucketTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
