package calendareventinstanceexceptionoccurrence

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceExceptionOccurrenceClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceExceptionOccurrenceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceExceptionOccurrenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstanceexceptionoccurrence", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceExceptionOccurrenceClient: %+v", err)
	}

	return &CalendarEventInstanceExceptionOccurrenceClient{
		Client: client,
	}, nil
}
