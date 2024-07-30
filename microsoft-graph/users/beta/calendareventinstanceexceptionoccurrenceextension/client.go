package calendareventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &CalendarEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
