package usercalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewClient: %+v", err)
	}

	return &UserCalendarViewClient{
		Client: client,
	}, nil
}
