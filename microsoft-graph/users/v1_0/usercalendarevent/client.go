package usercalendarevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarEventClient struct {
	Client *msgraph.Client
}

func NewUserCalendarEventClientWithBaseURI(api sdkEnv.Api) (*UserCalendarEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarEventClient: %+v", err)
	}

	return &UserCalendarEventClient{
		Client: client,
	}, nil
}
