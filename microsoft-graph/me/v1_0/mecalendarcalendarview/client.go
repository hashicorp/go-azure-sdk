package mecalendarcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewClient: %+v", err)
	}

	return &MeCalendarCalendarViewClient{
		Client: client,
	}, nil
}
