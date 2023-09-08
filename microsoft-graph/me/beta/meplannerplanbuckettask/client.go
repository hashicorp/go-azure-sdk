package meplannerplanbuckettask

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanBucketTaskClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanBucketTaskClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanBucketTaskClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplanbuckettask", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanBucketTaskClient: %+v", err)
	}

	return &MePlannerPlanBucketTaskClient{
		Client: client,
	}, nil
}
