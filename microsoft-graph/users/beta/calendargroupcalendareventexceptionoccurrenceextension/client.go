package calendargroupcalendareventexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &CalendarGroupCalendarEventExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
