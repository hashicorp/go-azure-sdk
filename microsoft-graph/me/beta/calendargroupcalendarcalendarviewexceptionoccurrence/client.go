package calendargroupcalendarcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
