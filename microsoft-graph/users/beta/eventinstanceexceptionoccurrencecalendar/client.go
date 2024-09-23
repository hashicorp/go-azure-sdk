package eventinstanceexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventInstanceExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewEventInstanceExceptionOccurrenceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*EventInstanceExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "eventinstanceexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EventInstanceExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &EventInstanceExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
