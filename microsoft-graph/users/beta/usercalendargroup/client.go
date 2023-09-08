package usercalendargroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupClient: %+v", err)
	}

	return &UserCalendarGroupClient{
		Client: client,
	}, nil
}
