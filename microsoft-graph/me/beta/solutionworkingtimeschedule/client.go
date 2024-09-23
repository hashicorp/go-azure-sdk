package solutionworkingtimeschedule

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SolutionWorkingTimeScheduleClient struct {
	Client *msgraph.Client
}

func NewSolutionWorkingTimeScheduleClientWithBaseURI(sdkApi sdkEnv.Api) (*SolutionWorkingTimeScheduleClient, error) {
	client, err := msgraph.NewClient(sdkApi, "solutionworkingtimeschedule", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SolutionWorkingTimeScheduleClient: %+v", err)
	}

	return &SolutionWorkingTimeScheduleClient{
		Client: client,
	}, nil
}
