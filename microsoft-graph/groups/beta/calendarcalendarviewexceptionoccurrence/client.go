package calendarcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
