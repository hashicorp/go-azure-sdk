package mecalendarviewinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarViewInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarViewInstanceCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarViewInstanceCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendarviewinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarViewInstanceCalendarClient: %+v", err)
	}

	return &MeCalendarViewInstanceCalendarClient{
		Client: client,
	}, nil
}
