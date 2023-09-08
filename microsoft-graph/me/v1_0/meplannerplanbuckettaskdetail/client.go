package meplannerplanbuckettaskdetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanBucketTaskDetailClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanBucketTaskDetailClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanBucketTaskDetailClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplanbuckettaskdetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanBucketTaskDetailClient: %+v", err)
	}

	return &MePlannerPlanBucketTaskDetailClient{
		Client: client,
	}, nil
}
