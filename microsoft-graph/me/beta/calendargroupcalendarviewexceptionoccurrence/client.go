package calendargroupcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarGroupCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
