package teamscheduletimeoffreason

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleTimeOffReasonClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleTimeOffReasonClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleTimeOffReasonClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduletimeoffreason", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleTimeOffReasonClient: %+v", err)
	}

	return &TeamScheduleTimeOffReasonClient{
		Client: client,
	}, nil
}
