package usercalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &UserCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
