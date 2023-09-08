package groupcalendarviewexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &GroupCalendarViewExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
