package calendareventinstancecalendar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarEventInstanceCalendarClient struct {
	Client *msgraph.Client
}

func NewCalendarEventInstanceCalendarClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarEventInstanceCalendarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendareventinstancecalendar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarEventInstanceCalendarClient: %+v", err)
	}

	return &CalendarEventInstanceCalendarClient{
		Client: client,
	}, nil
}
