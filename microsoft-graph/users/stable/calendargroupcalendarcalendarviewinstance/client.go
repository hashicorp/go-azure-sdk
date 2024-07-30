package calendargroupcalendarcalendarviewinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarCalendarViewInstanceClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarCalendarViewInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarCalendarViewInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarcalendarviewinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarCalendarViewInstanceClient: %+v", err)
	}

	return &CalendarGroupCalendarCalendarViewInstanceClient{
		Client: client,
	}, nil
}
