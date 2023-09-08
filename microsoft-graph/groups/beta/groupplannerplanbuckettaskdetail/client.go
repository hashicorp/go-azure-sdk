package groupplannerplanbuckettaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanBucketTaskDetailClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanBucketTaskDetailClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanBucketTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplanbuckettaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanBucketTaskDetailClient: %+v", err)
	}

	return &GroupPlannerPlanBucketTaskDetailClient{
		Client: client,
	}, nil
}
