package usercalendarcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &UserCalendarCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
