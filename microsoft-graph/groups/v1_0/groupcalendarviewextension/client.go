package groupcalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewExtensionClient: %+v", err)
	}

	return &GroupCalendarViewExtensionClient{
		Client: client,
	}, nil
}
