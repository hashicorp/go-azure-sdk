package usercalendarcalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarCalendarViewInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarCalendarViewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarCalendarViewInstanceClient: %+v", err)
	}

	return &UserCalendarCalendarViewInstanceClient{
		Client: client,
	}, nil
}
