package plannerplanbucket

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanBucketClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanBucketClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanBucketClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplanbucket", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanBucketClient: %+v", err)
	}

	return &PlannerPlanBucketClient{
		Client: client,
	}, nil
}
