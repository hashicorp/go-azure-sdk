package calendareventexceptionoccurrenceinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrenceinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceInstanceExtensionClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceInstanceExtensionClient{
		Client: client,
	}, nil
}
