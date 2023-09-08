package groupcalendareventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarEventInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarEventInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendareventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarEventInstanceExtensionClient: %+v", err)
	}

	return &GroupCalendarEventInstanceExtensionClient{
		Client: client,
	}, nil
}
