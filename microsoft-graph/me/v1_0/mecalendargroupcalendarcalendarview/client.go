package mecalendargroupcalendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewClient{
		Client: client,
	}, nil
}
