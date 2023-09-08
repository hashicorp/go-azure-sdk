package usercalendargroupcalendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
