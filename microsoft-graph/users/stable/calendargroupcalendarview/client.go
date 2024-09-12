package calendargroupcalendarview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupCalendarViewClient struct {
	Client *msgraph.Client
}

func NewCalendarGroupCalendarViewClientWithBaseURI(sdkApi sdkEnv.Api) (*CalendarGroupCalendarViewClient, error) {
	client, err := msgraph.NewClient(sdkApi, "calendargroupcalendarview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CalendarGroupCalendarViewClient: %+v", err)
	}

	return &CalendarGroupCalendarViewClient{
		Client: client,
	}, nil
}
