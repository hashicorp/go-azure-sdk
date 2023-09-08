package usercalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &UserCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
