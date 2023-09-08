package meeventexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeEventExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeEventExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeEventExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
