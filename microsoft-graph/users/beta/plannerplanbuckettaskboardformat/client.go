package plannerplanbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
