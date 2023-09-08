package usereventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewUserEventInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*UserEventInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usereventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserEventInstanceCalendarClient: %+v", err)
	}

	return &UserEventInstanceCalendarClient{
		Client: client,
	}, nil
}
