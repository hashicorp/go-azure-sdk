package mecalendargroupcalendarcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
