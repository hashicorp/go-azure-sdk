package userplannertaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPlannerTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewUserPlannerTaskProgressTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*UserPlannerTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userplannertaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserPlannerTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &UserPlannerTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
