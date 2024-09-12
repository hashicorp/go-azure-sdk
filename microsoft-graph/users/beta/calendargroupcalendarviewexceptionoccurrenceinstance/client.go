package calendargroupcalendarviewexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewExceptionOccurrenceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &CalendarGroupCalendarViewExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
