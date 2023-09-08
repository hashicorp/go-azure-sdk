package groupcalendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventCalendarClient: %+v", err)
	}

	return &GroupCalendarEventCalendarClient{
		Client: client,
	}, nil
}
