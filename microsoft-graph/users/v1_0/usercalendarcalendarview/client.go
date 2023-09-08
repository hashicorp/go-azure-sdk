package usercalendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewClient: %+v", err)
	}

	return &UserCalendarCalendarViewClient{
		Client: client,
	}, nil
}
