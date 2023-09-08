package mecalendarcalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarcalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &MeCalendarCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
