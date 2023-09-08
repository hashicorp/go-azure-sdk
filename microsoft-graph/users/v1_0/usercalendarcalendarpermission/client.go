package usercalendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarPermissionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarPermissionClient: %+v", err)
	}

	return &UserCalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
