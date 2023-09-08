package groupcalendarcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
