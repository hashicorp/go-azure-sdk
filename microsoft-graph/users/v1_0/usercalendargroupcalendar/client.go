package usercalendargroupcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarClient: %+v", err)
	}

	return &UserCalendarGroupCalendarClient{
		Client: client,
	}, nil
}
