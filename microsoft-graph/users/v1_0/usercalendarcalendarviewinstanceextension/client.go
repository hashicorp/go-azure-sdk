package usercalendarcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewInstanceExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &UserCalendarCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
