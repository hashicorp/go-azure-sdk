package usereventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventCalendarClient struct {
	Client *msgraph.Client
}

func NewUserEventCalendarClientWithBaseURI(api sdkEnv.Api) (*UserEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventCalendarClient: %+v", err)
	}

	return &UserEventCalendarClient{
		Client: client,
	}, nil
}
