package calendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarEventCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventCalendarClient: %+v", err)
	}

	return &CalendarEventCalendarClient{
		Client: client,
	}, nil
}
