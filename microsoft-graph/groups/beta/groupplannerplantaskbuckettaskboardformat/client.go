package groupplannerplantaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplantaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &GroupPlannerPlanTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
