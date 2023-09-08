package usercalendargroupcalendareventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
