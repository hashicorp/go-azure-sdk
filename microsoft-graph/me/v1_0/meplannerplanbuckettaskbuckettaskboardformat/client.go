package meplannerplanbuckettaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanBucketTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanBucketTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanBucketTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplanbuckettaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanBucketTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &MePlannerPlanBucketTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
