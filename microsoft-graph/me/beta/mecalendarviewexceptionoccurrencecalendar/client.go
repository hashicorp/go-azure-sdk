package mecalendarviewexceptionoccurrencecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewExceptionOccurrenceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewExceptionOccurrenceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewExceptionOccurrenceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewexceptionoccurrencecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewExceptionOccurrenceCalendarClient: %+v", err)
	}

	return &MeCalendarViewExceptionOccurrenceCalendarClient{
		Client: client,
	}, nil
}
