package groupcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarClient: %+v", err)
	}

	return &GroupCalendarClient{
		Client: client,
	}, nil
}
