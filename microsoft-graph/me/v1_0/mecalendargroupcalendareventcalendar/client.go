package mecalendargroupcalendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarEventCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarEventCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarEventCalendarClient{
		Client: client,
	}, nil
}
