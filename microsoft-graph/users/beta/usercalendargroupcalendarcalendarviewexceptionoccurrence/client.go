package usercalendargroupcalendarcalendarviewexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewExceptionOccurrenceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewExceptionOccurrenceClient{
		Client: client,
	}, nil
}
