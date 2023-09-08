package meplannerplantaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MePlannerPlanTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewMePlannerPlanTaskAssignedToTaskBoardFormatClientWithBaseURI(api sdkEnv.Api) (*MePlannerPlanTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meplannerplantaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MePlannerPlanTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &MePlannerPlanTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
