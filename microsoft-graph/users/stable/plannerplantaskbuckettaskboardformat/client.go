package plannerplantaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplantaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
