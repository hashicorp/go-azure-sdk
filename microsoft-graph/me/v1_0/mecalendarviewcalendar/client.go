package mecalendarviewcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewCalendarClient: %+v", err)
	}

	return &MeCalendarViewCalendarClient{
		Client: client,
	}, nil
}
