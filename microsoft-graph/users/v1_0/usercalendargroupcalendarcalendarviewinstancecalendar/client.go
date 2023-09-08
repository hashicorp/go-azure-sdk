package usercalendargroupcalendarcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
