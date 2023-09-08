package usercalendarcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
