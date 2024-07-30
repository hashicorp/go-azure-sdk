package calendargroupcalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarClient: %+v", err)
	}

	return &CalendarGroupCalendarClient{
		Client: client,
	}, nil
}
