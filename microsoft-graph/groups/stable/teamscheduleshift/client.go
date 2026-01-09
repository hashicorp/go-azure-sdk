package teamscheduleshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleShiftClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleShiftClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleShiftClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduleshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleShiftClient: %+v", err)
	}

	return &TeamScheduleShiftClient{
		Client: client,
	}, nil
}
