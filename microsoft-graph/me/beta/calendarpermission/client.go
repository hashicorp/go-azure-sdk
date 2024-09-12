package calendarpermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermissionClient struct {
	Client *msgraph.Client
}

func NewCalendarPermissionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarPermissionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendarpermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarPermissionClient: %+v", err)
	}

	return &CalendarPermissionClient{
		Client: client,
	}, nil
}
