package calendareventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &CalendarEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
