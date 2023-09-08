package usercalendargroupcalendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventCalendarClient{
		Client: client,
	}, nil
}
