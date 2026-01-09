package teamscheduletimesoff

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleTimesOffClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleTimesOffClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleTimesOffClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduletimesoff", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleTimesOffClient: %+v", err)
	}

	return &TeamScheduleTimesOffClient{
		Client: client,
	}, nil
}
