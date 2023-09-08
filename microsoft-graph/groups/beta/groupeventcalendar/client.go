package groupeventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupEventCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventCalendarClient: %+v", err)
	}

	return &GroupEventCalendarClient{
		Client: client,
	}, nil
}
