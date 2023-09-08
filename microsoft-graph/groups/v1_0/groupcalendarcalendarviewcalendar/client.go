package groupcalendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewCalendarClient: %+v", err)
	}

	return &GroupCalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
