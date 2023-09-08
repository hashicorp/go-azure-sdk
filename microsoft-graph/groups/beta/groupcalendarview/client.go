package groupcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupCalendarViewClient struct {
	Client *msgraph.Client
}

func NewGroupCalendarViewClientWithBaseURI(api sdkEnv.Api) (*GroupCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupCalendarViewClient: %+v", err)
	}

	return &GroupCalendarViewClient{
		Client: client,
	}, nil
}
