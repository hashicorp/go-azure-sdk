package groupcalendarviewinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupCalendarViewInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
