package groupcalendarcalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
