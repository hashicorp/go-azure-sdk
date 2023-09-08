package mecalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarviewexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
