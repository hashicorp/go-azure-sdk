package groupcalendareventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventExtensionClient: %+v", err)
	}

	return &GroupCalendarEventExtensionClient{
		Client: client,
	}, nil
}
