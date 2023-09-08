package usercalendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventCalendarClient: %+v", err)
	}

	return &UserCalendarEventCalendarClient{
		Client: client,
	}, nil
}
