package groupplannerplanbuckettaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanBucketTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanBucketTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplanbuckettaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanBucketTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &GroupPlannerPlanBucketTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
