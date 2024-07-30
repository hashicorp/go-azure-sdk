package plannerplanbuckettaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbuckettaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanBucketTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
