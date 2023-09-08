package usercalendargroupcalendareventextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventExtensionClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventExtensionClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventExtensionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventExtensionClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventExtensionClient{
		Client: client,
	}, nil
}
