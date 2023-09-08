package usercalendargroupcalendarevent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCalendarGroupCalendarEventClient struct {
	Client *msgraph.Client
}

func NewUserCalendarGroupCalendarEventClientWithBaseURI(api sdkEnv.Api) (*UserCalendarGroupCalendarEventClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercalendargroupcalendarevent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCalendarGroupCalendarEventClient: %+v", err)
	}

	return &UserCalendarGroupCalendarEventClient{
		Client: client,
	}, nil
}
