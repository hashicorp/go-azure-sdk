package usercalendareventinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceExtensionClient: %+v", err)
	}

	return &UserCalendarEventInstanceExtensionClient{
		Client: client,
	}, nil
}
