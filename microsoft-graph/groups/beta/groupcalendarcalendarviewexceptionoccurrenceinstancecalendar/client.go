package groupcalendarcalendarviewexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &GroupCalendarCalendarViewExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
