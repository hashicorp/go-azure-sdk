package calendargroupcalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
