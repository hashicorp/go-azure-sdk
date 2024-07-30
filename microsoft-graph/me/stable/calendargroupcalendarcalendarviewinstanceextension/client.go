package calendargroupcalendarcalendarviewinstanceextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewInstanceExtensionClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewInstanceExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewInstanceExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewinstanceextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewInstanceExtensionClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewInstanceExtensionClient{
		Client: client,
	}, nil
}
