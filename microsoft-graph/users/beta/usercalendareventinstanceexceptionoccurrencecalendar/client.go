package usercalendareventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &UserCalendarEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
