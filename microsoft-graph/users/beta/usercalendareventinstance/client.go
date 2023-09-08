package usercalendareventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendareventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventInstanceClient: %+v", err)
	}

	return &UserCalendarEventInstanceClient{
		Client: client,
	}, nil
}
