package groupcalendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewClient: %+v", err)
	}

	return &GroupCalendarCalendarViewClient{
		Client: client,
	}, nil
}
