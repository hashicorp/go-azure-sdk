package userplannerplantaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerPlanTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerPlanTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannerplantaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerPlanTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerPlanTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
