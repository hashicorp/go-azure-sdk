package plannerplanbuckettaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketTaskDetailClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketTaskDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketTaskDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbuckettaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketTaskDetailClient: %+v", err)
	}

	return &PlannerPlanBucketTaskDetailClient{
		Client: client,
	}, nil
}
