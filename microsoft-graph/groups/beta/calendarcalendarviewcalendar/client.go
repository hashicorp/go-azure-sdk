package calendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewCalendarClient: %+v", err)
	}

	return &CalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
