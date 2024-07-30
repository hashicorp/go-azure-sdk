package calendargroupcalendarcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
