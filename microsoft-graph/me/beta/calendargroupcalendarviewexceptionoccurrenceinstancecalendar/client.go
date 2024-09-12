package calendargroupcalendarviewexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarViewExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
