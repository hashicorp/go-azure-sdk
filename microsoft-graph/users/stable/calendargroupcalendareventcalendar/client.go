package calendargroupcalendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarEventCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarEventCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarEventCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarEventCalendarClient{
		Client: client,
	}, nil
}
