package userplannerplantaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplantaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerPlanTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
