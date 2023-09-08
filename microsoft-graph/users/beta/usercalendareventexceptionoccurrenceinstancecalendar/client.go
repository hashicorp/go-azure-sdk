package usercalendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &UserCalendarEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
