package calendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarViewInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarViewInstanceCalendarClient: %+v", err)
	}

	return &CalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
