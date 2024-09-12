package teamscheduleoffershiftrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamScheduleOfferShiftRequestClient struct {
	Client *msgraph.Client
}

func NewTeamScheduleOfferShiftRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*TeamScheduleOfferShiftRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "teamscheduleoffershiftrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TeamScheduleOfferShiftRequestClient: %+v", err)
	}

	return &TeamScheduleOfferShiftRequestClient{
		Client: client,
	}, nil
}
