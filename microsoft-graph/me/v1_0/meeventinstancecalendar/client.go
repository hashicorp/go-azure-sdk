package meeventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeEventInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeEventInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meeventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeEventInstanceCalendarClient: %+v", err)
	}

	return &MeEventInstanceCalendarClient{
		Client: client,
	}, nil
}
