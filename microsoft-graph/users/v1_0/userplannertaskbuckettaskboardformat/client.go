package userplannertaskbuckettaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerTaskBucketTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerTaskBucketTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerTaskBucketTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannertaskbuckettaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerTaskBucketTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerTaskBucketTaskBoardFormatClient{
		Client: client,
	}, nil
}
