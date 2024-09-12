package calendarcalendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarCalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewCalendarCalendarPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarCalendarPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarcalendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarCalendarPermissionClient: %+v", err)
	}

	return &CalendarCalendarPermissionClient{
		Client: client,
	}, nil
}
