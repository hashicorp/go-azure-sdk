package mecalendargroupcalendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarCalendarPermissionClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarCalendarPermissionClient: %+v", err)
	}

	return &MeCalendarGroupCalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
