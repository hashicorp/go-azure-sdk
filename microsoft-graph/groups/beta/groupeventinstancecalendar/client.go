package groupeventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupEventInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupEventInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupeventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupEventInstanceCalendarClient: %+v", err)
	}

	return &GroupEventInstanceCalendarClient{
		Client: client,
	}, nil
}
