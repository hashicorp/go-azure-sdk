package calendargroupcalendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
