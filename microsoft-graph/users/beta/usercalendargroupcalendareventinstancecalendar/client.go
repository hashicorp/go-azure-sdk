package usercalendargroupcalendareventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventInstanceCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventInstanceCalendarClient{
		Client: client,
	}, nil
}
