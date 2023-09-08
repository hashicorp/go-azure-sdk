package mecalendareventcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeCalendarEventCalendarClient struct {
	Client *msgraph.Client
}

func NewMeCalendarEventCalendarClientWithBaseURI(api sdkEnv.Api) (*MeCalendarEventCalendarClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mecalendareventcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeCalendarEventCalendarClient: %+v", err)
	}

	return &MeCalendarEventCalendarClient{
		Client: client,
	}, nil
}
