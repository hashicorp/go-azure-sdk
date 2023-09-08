package groupcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewCalendarClient: %+v", err)
	}

	return &GroupCalendarViewCalendarClient{
		Client: client,
	}, nil
}
