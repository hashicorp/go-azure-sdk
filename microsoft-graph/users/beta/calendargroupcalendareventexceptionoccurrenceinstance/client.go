package calendargroupcalendareventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &CalendarGroupCalendarEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
