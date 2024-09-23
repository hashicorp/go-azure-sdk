package teamscheduleopenshiftchangerequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleOpenShiftChangeRequestClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleOpenShiftChangeRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleOpenShiftChangeRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduleopenshiftchangerequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleOpenShiftChangeRequestClient: %+v", err)
	}

	return &TeamScheduleOpenShiftChangeRequestClient{
		Client: client,
	}, nil
}
