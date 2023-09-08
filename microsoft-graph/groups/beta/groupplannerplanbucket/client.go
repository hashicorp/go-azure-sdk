package groupplannerplanbucket

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanBucketClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanBucketClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanBucketClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplanbucket", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanBucketClient: %+v", err)
	}

	return &GroupPlannerPlanBucketClient{
		Client: client,
	}, nil
}
