package usercalendareventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
