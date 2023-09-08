package groupcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &GroupCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
