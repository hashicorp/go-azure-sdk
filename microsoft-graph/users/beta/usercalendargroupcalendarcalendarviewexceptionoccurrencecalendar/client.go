package usercalendargroupcalendarcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
