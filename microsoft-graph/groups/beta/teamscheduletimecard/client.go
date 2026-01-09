package teamscheduletimecard

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleTimeCardClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleTimeCardClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleTimeCardClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduletimecard", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleTimeCardClient: %+v", err)
	}

	return &TeamScheduleTimeCardClient{
		Client: client,
	}, nil
}
