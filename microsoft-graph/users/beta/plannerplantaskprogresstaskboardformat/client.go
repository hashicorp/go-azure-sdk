package plannerplantaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerPlanTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerPlanTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannerplantaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerPlanTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &PlannerPlanTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
