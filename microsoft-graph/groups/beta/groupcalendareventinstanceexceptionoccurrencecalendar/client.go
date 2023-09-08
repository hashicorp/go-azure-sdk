package groupcalendareventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &GroupCalendarEventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
