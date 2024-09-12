package calendargroupcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
