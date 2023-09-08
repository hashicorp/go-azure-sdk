package groupcalendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExceptionOccurrenceInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExceptionOccurrenceInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExceptionOccurrenceInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventexceptionoccurrenceinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExceptionOccurrenceInstanceCalendarClient: %+v", err)
	}

	return &GroupCalendarEventExceptionOccurrenceInstanceCalendarClient{
		Client: client,
	}, nil
}
