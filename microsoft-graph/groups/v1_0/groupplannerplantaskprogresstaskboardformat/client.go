package groupplannerplantaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPlannerPlanTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewGroupPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*GroupPlannerPlanTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupplannerplantaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPlannerPlanTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &GroupPlannerPlanTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
