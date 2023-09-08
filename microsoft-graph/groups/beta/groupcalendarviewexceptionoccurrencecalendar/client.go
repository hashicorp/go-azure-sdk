package groupcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
