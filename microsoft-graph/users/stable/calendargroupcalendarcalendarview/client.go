package calendargroupcalendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewClient{
		Client: client,
	}, nil
}
