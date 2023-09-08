package groupcalendareventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
