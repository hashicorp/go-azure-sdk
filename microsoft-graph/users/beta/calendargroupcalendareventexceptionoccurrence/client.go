package calendargroupcalendareventexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarGroupCalendarEventExceptionOccurrenceClient{
		Client: client,
	}, nil
}
