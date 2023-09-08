package groupteamscheduleoffershiftrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTeamScheduleOfferShiftRequestClient struct {
	Client *msgraph.Client
}

func NewGroupTeamScheduleOfferShiftRequestClientWithBaseURI(api sdkEnv.Api) (*GroupTeamScheduleOfferShiftRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupteamscheduleoffershiftrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTeamScheduleOfferShiftRequestClient: %+v", err)
	}

	return &GroupTeamScheduleOfferShiftRequestClient{
		Client: client,
	}, nil
}
