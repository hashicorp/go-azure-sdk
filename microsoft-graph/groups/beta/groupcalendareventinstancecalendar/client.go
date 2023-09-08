package groupcalendareventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceCalendarClient: %+v", err)
	}

	return &GroupCalendarEventInstanceCalendarClient{
		Client: client,
	}, nil
}
