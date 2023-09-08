package usercalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
