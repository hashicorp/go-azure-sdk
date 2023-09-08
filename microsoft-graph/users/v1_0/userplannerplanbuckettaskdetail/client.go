package userplannerplanbuckettaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanBucketTaskDetailClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanBucketTaskDetailClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanBucketTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplanbuckettaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanBucketTaskDetailClient: %+v", err)
	}

	return &UserPlannerPlanBucketTaskDetailClient{
		Client: client,
	}, nil
}
