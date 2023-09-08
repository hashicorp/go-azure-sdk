package userplannerplanbucket

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanBucketClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanBucketClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanBucketClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplanbucket", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanBucketClient: %+v", err)
	}

	return &UserPlannerPlanBucketClient{
		Client: client,
	}, nil
}
