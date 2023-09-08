package usercalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarViewInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarViewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarViewInstanceClient: %+v", err)
	}

	return &UserCalendarViewInstanceClient{
		Client: client,
	}, nil
}
