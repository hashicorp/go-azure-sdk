package plannertaskprogresstaskboardformat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTaskProgressTaskBoardFormatClient struct {
	Client *msgraph.Client
}

func NewPlannerTaskProgressTaskBoardFormatClientWithBaseURI(sdkApi sdkEnv.Api) (*PlannerTaskProgressTaskBoardFormatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "plannertaskprogresstaskboardformat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlannerTaskProgressTaskBoardFormatClient: %+v", err)
	}

	return &PlannerTaskProgressTaskBoardFormatClient{
		Client: client,
	}, nil
}
