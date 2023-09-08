package mecalendargroupcalendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
