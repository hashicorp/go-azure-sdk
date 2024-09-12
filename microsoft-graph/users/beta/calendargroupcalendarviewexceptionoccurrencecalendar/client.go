package calendargroupcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
