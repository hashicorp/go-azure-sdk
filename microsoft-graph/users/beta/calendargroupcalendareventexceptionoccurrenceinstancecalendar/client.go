package calendargroupcalendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
