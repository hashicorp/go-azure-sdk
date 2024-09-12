package calendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarViewExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &CalendarViewExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
