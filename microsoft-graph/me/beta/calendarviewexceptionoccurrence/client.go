package calendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarViewExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
