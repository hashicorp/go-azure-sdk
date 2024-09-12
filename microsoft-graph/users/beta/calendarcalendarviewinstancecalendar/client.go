package calendarcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarViewInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &CalendarCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
