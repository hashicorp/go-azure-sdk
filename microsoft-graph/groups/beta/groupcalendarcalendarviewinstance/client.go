package groupcalendarcalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarCalendarViewInstanceClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarCalendarViewInstanceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarCalendarViewInstanceClient: %+v", err)
	}

	return &GroupCalendarCalendarViewInstanceClient{
		Client: client,
	}, nil
}
