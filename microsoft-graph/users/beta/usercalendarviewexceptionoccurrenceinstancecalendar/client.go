package usercalendarviewexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &UserCalendarViewExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
