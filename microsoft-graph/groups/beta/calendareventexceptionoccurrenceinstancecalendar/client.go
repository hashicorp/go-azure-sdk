package calendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
