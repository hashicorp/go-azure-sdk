package userplannerplanbuckettask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanBucketTaskClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanBucketTaskClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanBucketTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplanbuckettask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanBucketTaskClient: %+v", err)
	}

	return &UserPlannerPlanBucketTaskClient{
		Client: client,
	}, nil
}
