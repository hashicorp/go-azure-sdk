package calendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &CalendarViewInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
