package groupcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &GroupCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
