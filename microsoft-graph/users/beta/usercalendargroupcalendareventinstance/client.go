package usercalendargroupcalendareventinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventInstanceClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventInstanceClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendareventinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventInstanceClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventInstanceClient{
		Client: client,
	}, nil
}
