package mecalendarcalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
