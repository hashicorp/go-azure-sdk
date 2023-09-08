package groupplannerplanbuckettaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanBucketTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanBucketTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplanbuckettaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanBucketTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &GroupPlannerPlanBucketTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
