package plannertaskassignedtotaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskAssignedToTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerTaskAssignedToTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerTaskAssignedToTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannertaskassignedtotaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerTaskAssignedToTaskBoardFormatClient: %+v", err)
	}

	return &PlannerTaskAssignedToTaskBoardFormatClient{
		Client: client,
	}, nil
}
