package calendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
