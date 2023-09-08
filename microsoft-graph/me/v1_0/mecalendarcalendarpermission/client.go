package mecalendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarPermissionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarPermissionClient: %+v", err)
	}

	return &MeCalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
