package mecalendarcalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewCalendarClient: %+v", err)
	}

	return &MeCalendarCalendarViewCalendarClient{
		Client: client,
	}, nil
}
