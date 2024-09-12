package calendargroupcalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarGroupCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
