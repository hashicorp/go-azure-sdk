package teamscheduleopenshift

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleOpenShiftClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleOpenShiftClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleOpenShiftClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduleopenshift", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleOpenShiftClient: %+v", err)
	}

	return &TeamScheduleOpenShiftClient{
		Client: client,
	}, nil
}
