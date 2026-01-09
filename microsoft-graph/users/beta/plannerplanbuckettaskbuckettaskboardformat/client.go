package plannerplanbuckettaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbuckettaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanBucketTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
