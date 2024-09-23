package calendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
