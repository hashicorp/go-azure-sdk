package calendargroupcalendareventinstanceexceptionoccurrenceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventinstanceexceptionoccurrenceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient: %+v", err)
	}

	return &CalendarGroupCalendarEventInstanceExceptionOccurrenceExtensionClient{
		Client: client,
	}, nil
}
