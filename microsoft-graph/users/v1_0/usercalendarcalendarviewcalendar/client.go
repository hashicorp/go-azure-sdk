package usercalendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewCalendarClient: %+v", err)
	}

	return &UserCalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
