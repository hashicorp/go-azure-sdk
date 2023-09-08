package usercalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
