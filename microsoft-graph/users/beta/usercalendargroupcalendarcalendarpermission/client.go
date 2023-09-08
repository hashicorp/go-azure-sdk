package usercalendargroupcalendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarCalendarPermissionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarCalendarPermissionClient: %+v", err)
	}

	return &UserCalendarGroupCalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
