package calendargroupcalendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarPermissionClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
