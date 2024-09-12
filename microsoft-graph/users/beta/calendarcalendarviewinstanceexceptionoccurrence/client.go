package calendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarCalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
