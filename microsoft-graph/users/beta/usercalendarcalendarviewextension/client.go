package usercalendarcalendarviewextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewExtensionClient: %+v", err)
	}

	return &UserCalendarCalendarViewExtensionClient{
		Client: client,
	}, nil
}
