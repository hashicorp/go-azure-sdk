package plannerplanbuckettask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketTaskClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketTaskClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketTaskClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbuckettask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketTaskClient: %+v", err)
	}

	return &PlannerPlanBucketTaskClient{
		Client: client,
	}, nil
}
