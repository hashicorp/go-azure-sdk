package userplannerplanbuckettaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanBucketTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanBucketTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplanbuckettaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanBucketTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerPlanBucketTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
