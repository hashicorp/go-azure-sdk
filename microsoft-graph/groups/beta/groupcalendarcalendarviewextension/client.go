package groupcalendarcalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewExtensionClient: %+v", err)
	}

	return &GroupCalendarCalendarViewExtensionClient{
		Client: client,
	}, nil
}
