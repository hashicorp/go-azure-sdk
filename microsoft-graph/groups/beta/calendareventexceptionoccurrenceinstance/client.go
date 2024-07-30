package calendareventexceptionoccurrenceinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventExceptionOccurrenceInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarEventExceptionOccurrenceInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventExceptionOccurrenceInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventexceptionoccurrenceinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventExceptionOccurrenceInstanceClient: %+v", err)
	}

	return &CalendarEventExceptionOccurrenceInstanceClient{
		Client: client,
	}, nil
}
