package usercalendargroupcalendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewClient{
		Client: client,
	}, nil
}
