package meplannerplanbuckettaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanBucketTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanBucketTaskProgressTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanBucketTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplanbuckettaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanBucketTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &MePlannerPlanBucketTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
