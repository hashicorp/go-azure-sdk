package usercalendareventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventExtensionClient: %+v", err)
	}

	return &UserCalendarEventExtensionClient{
		Client: client,
	}, nil
}
