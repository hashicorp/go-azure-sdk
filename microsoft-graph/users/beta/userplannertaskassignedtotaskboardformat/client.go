package userplannertaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerTaskAssignedToTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannertaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
