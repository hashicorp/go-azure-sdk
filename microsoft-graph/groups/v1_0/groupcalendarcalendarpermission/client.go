package groupcalendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarPermissionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarPermissionClient: %+v", err)
	}

	return &GroupCalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
