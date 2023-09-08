package usercalendargroupcalendarcalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarViewExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarViewExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarViewExtensionClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarViewExtensionClient{
		Client: client,
	}, nil
}
