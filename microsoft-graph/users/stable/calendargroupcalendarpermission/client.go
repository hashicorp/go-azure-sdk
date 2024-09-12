package calendargroupcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarPermissionClient: %+v", err)
	}

	return &CalendarGroupCalendarPermissionClient{
		Client: client,
	}, nil
}
