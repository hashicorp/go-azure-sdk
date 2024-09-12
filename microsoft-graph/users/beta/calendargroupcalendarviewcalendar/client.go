package calendargroupcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarViewCalendarClient{
		Client: client,
	}, nil
}
