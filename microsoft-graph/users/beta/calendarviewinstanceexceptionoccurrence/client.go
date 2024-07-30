package calendarviewinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarViewInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
