package mecalendargroupcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarGroupCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarGroupCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarGroupCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendargroupcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarGroupCalendarClient: %+v", err)
	}

	return &MeCalendarGroupCalendarClient{
		Client: client,
	}, nil
}
